package controllers

import (	
	"fmt"
	"log"
	"net/http"
	"github.com/igorvinnicius/lenslocked-go-web/views"
	"github.com/igorvinnicius/lenslocked-go-web/models"
)

func NewGalleries(galleryService models.GalleryService) *Galleries {
	return &Galleries{
		New: views.NewView("bootstrap", "galleries/new"),		
		GalleryService : galleryService,
	}
}

type Galleries struct{
	New *views.View	
	GalleryService models.GalleryService
}

type GalleryForm struct {
	Title string `schema:"title"`	
}

func (g *Galleries) Create(w http.ResponseWriter, r *http.Request) {
		
	var vd views.Data
	var form GalleryForm

	if err := parseForm(r, &form); err != nil {
		log.Println(err)
		vd.SetAlert(err)
		g.New.Render(w, vd)
		return
	}

	gallery := models.Gallery {
		Title: form.Title,
	}
	
	if err := g.GalleryService.Create(&gallery); err != nil {		
		vd.SetAlert(err)
		g.New.Render(w, vd)
		return		
	}

	fmt.Fprintln(w, gallery)
}