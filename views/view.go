package views

import (
	"html/template"
	"path/filepath"
	"net/http"
)

var (
	LayoutDirectory = "views/layouts/"
	TemplateDir = "views/"
	LayoutExtension = ".gohtml"
)

type View struct{
	Template *template.Template
	Layout string
}

func NewView(layout string, files ...string) *View {
	
	addTemplatePath(files)
	addTemplateExtension(files)

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
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

func layoutFiles() []string{
	
	files, err := filepath.Glob(LayoutDirectory + "*" + LayoutExtension)
	
	if err != nil{
		panic(err)
	}
	
	return files
}

func addTemplatePath(files []string) {
	for i,f := range files {
		files[i] = TemplateDir + f
	}
}

func addTemplateExtension(files []string) {
	for i,f := range files {
		files[i] = f + LayoutExtension
	}
}


