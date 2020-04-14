package models

import "github.com/jinzhu/gorm"

type Gallery struct {
	gorm.Model
	UserId uint `gorm:"not null;index"`
	Title string `gorm:"not null"`
}