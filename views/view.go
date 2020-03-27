package views

import (
	"html/template"
	"path/filepath"
	"net/http"
)

var (
	LayoutDirectory = "views/layouts/"
	LayoutExtension = ".gohtml"
)

type View struct{
	Template *template.Template
	Layout string
}

func NewView(layout string, files ...string) *View {
	
	files = append(files, layoutFiles()...)
	
	t, err := template.ParseFiles(files...)
	
	if err != nil{
		panic(err)
	}

	return &View{
		Template: t,
		Layout: layout,
	}
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func layoutFiles() []string{
	
	files, err := filepath.Glob(LayoutDirectory + "*" + LayoutExtension)
	
	if err != nil{
		panic(err)
	}
	
	return files
}

