package controllers

import(
	"github.com/igorvinnicius/lenslocked-go-web/views"
)

type Static struct {
	HomeView *views.View
	ContactView *views.View
}

func NewStatic() *Static {
	return &Static{
		HomeView: views.NewView("bootstrap", "static/home"),
		ContactView: views.NewView("bootstrap", "static/contact"),
	}
}