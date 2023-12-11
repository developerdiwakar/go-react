package models

import "gorm.io/gorm"

type Business struct {
	gorm.Model
	Name        string
	Description string
	Image       string
	UserID      int
	User        User
}
