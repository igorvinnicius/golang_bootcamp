package main

import (
	"fmt"
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
}

func main() {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlinfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.AutoMigrate(&User{})
}

