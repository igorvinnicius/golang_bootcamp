package views

import (
	"html/template"
	"path/filepath"
	"net/http"
	"bytes"
	"io"
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

func (v *View) Render(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html")

	switch data.(type){
		case Data:
			//do nothing
		default: 
			data = Data {
				Yield: data,
			}
	}

	var buf bytes.Buffer

	if err := v.Template.ExecuteTemplate(&buf, v.Layout, data); err != nil {

		http.Error(w, "Something went wrong!", http.StatusInternalServerError)
		return
	}

	io.Copy(w, &buf)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.Render(w, nil)	
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


