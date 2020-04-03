package models

import(
	"fmt"
	"testing"
)

func testingUserService() (*UserService, error) {

	const(
		host = "localhost"
		port = 5432
		user = "postgres"
		password = "0000"
		dbname = "lenslocked_test"
	)
	
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	us, err := NewUserService(psqlinfo)

	if err != nil {
		return nil, err
	}

	us.db.LogMode(false)
	us.DestructiveReset()

	return us, nil	
}

func TestCreateUser(t *testing.T) {
	us, err := testingUserService()
	
	if err != nil {
		t.Fatal(err)
	}

	user := User {
		Name: "Jon Snow",
		Email: "jonsnow@starks.com",
	}

	err = us.Create(&user)

	if err != nil {
		t.Fatal(err)
	}

	if user.ID == 0 {
		t.Errorf("Expected ID > 0. Received %d", user.ID)
	}
}