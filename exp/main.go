package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const(
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "0000"
	dbname = "lenslocked_dev"
)

type User struct {
	gorm.Model
	Name string
	Email string `gorm:"not null;unique_index"`
	Color string
}

func main() {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlinfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.LogMode(true)
	db.AutoMigrate(&User{})

	name, email, color := getInfo()

	u := User{
		Name : name,
		Email : email,
		Color: color,
	}

	if err = db.Create(&u).Error; err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", u)
}

func getInfo() (name, email, color string) {
	
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("What's your name?")
	
	name, _ = reader.ReadString('\n')

	fmt.Println("What's your email address?")
	
	email, _ = reader.ReadString('\n')

	fmt.Println("What's your favorit color?")
	
	color, _ = reader.ReadString('\n')
	
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	color = strings.TrimSpace(color)
	
	return name, email, color
}

