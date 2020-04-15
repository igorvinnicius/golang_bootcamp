package controllers

import (	
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