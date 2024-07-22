package models

import "gorm.io/gorm"

type BlogPost struct {
	gorm.Model
	Title string
	Body  string
}
