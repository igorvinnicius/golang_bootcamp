package models

import(
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"golang.org/x/crypto/bcrypt"

	"github.com/igorvinnicius/lenslocked-go-web/hash"
	"github.com/igorvinnicius/lenslocked-go-web/rand"
)

var(
	ErrNotFound = errors.New("models: resource not found")
	ErrInvalidID = errors.New("models: ID must me > 0")	
	ErrInvalidPassword = errors.New("models: incorrect password provided")
)

const userPwPepper = "secret-random-string"
const hmacSecretKey = "secret-hmac-key"

type UserDB interface {
	ById(id uint) (*User, error)
	ByEmail(email string) (*User, error)
	ByRemember(token string) (*User, error)

	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error

	Close() error

	AutoMigrate() error
	DestructiveReset() error
}

type User struct {
	gorm.Model
	Name string
	Email string `gorm:"not null;unique_index"`
	Password string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
	Remember string `gorm:"-"`
	RememberHash string `gorm:"not null;unique_index"`
}

func NewUserService(connectionInfo string) (*UserService, error) {
	
	ug, err := newUserGorm(connectionInfo)
	if err != nil {
		return nil, err
	}

	return &UserService{
		UserDB : &userValidator {
			UserDB: ug,
		},
	}, nil
}

func newUserGorm(connectionInfo string) (*userGorm, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	hmac := hash.NewHMAC(hmacSecretKey)

	return &userGorm{
		db: db,
		hmac: hmac,
	}, nil
}

type UserService struct {
	UserDB
}

type userValidator struct {
	UserDB
}

var _ UserDB = &userGorm{}

type userGorm struct {
	db *gorm.DB
	hmac hash.HMAC
}

func (ug *userGorm) ById(id uint) (*User, error) {
	var user User
	db := ug.db.Where("id = ?", id)
	err := first(db, &user)
	return &user, err
}

func (ug *userGorm) ByEmail(email string) (*User, error) {
	var user User
	db := ug.db.Where("email = ?", email)
	err := first(db, &user)
	return &user, err
}

func (ug *userGorm) ByRemember(token string) (*User, error) {
	var user User
	rememberHash := ug.hmac.Hash(token)	
	db := ug.db.Where("remember_hash = ?", rememberHash)
	err := first(db, &user)
	return &user, err
}

func first(db *gorm.DB, dest interface{}) error {
	
	err := db.First(dest).Error
	
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}

	return err
}

func (us *UserService) Authenticate(email, password string) (*User, error) {
	
	foundUser, err := us.ByEmail(email)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password + userPwPepper))

	if err != nil {
		switch err {
			case bcrypt.ErrMismatchedHashAndPassword:
				return nil, ErrInvalidPassword	
			default:
				return nil, err
		}
	}

	return foundUser, nil
}

func (ug *userGorm) Create(user *User) error {

	passwordBytes := []byte(user.Password + userPwPepper)

	hashedBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.PasswordHash = string(hashedBytes)

	if user.Remember == "" {
		
		token, err := rand.RememberToken()
		if err != nil {
			return err
		}

		user.Remember = token		
	}

	user.RememberHash = ug.hmac.Hash(user.Remember)

	return ug.db.Create(user).Error
}

func (ug *userGorm) Update(user *User) error {

	if user.Remember != "" {
		user.RememberHash = ug.hmac.Hash(user.Remember)
	}

	return ug.db.Save(user).Error
}

func (ug *userGorm) Delete(id uint) error {
	
	if id == 0 {
		return ErrInvalidID
	}

	user := User{Model: gorm.Model{ID: id}}
	
	return ug.db.Delete(&user).Error
}


func (ug *userGorm) Close() error {
	return ug.db.Close()
}

func (ug *userGorm) DestructiveReset() error {
	
	if err := ug.db.DropTableIfExists(&User{}).Error; err != nil {
		return err
	}

	return ug.AutoMigrate()
}

func (ug *userGorm) AutoMigrate() error {
	
	if err := ug.db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}

	return nil
}