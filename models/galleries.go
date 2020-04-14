package models

import "github.com/jinzhu/gorm"

type Gallery struct {
	gorm.Model
	UserId uint `gorm:"not null;index"`
	Title string `gorm:"not null"`
}

type GalleryService interface {
	GalleryDB
}

type GalleryDB interface {

	Create(gallery *Gallery) error

}

type galleryGorm struct {
	db *gorm.DB
}