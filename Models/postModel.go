package models

import (
	"time"

	"gorm.io/gorm"
)

type BlogPost struct {
	gorm.Model
	Title      string
	Body       string
	Category   string
	Author     string
	DatePosted time.Time
}
