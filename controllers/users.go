package controllers

import (
	"fmt"
	"net/http"
	"github.com/igorvinnicius/lenslocked-go-web/views"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct{
	NewView *views.View
}

func (u *Users) New(w http.ResponseWriter, r *http.Request){
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "This is a temporary response")
}