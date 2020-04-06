package models

import(
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var(
	ErrNotFound = errors.New("models: resource not found")
	ErrInvalidID = errors.New("models: ID must me > 0")
)

type User struct {
	gorm.Model
	Name string
	Email string `gorm:"not null;unique_index"`
}

func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}

	return &UserService{
		db: db,
	}, nil
}

type UserService struct {
	db *gorm.DB
}

func (us *UserService) ById(id uint) (*User, error) {
	var user User
	db := us.db.Where("id = ?", id)
	err := first(db, &user)
	return &user, err
}

func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email = ?", email)
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

func (us *UserService) Create(user *User) error {
	return us.db.Create(user).Error
}

func (us *UserService) Update(user *User) error {
	return us.db.Save(user).Error
}

func (us *UserService) Delete(id uint) error {
	
	if id == 0 {
		return ErrInvalidID
	}

	user := User{Model: gorm.Model{ID: id}}
	return us.db.Delete(&user).Error
}


func (us *UserService) Close() error {
	return us.db.Close()
}

func (us *UserService) DestructiveReset() {
	us.db.DropTableIfExists(&User{})
	us.db.AutoMigrate(&User{})
}