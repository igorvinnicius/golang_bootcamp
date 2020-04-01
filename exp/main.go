package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const(
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "0000"
	dbname = "lenslocked_dev"
)

func main() {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var id int

	err = db.QueryRow(`
		INSERT INTO users(name, email)
		values($1,$2)
		RETURNING id`,
		"Jon Snow", "snow@email.com").Scan(&id)

	if err != nil {
		panic(err)
	}

	fmt.Println("Id is...", id)

}

