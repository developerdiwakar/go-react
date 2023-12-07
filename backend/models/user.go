package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Name         string
	MobileNumber string `gorm:"uniqueIndex;size:15"`
	Email        string `gorm:"uniqueIndex;size:30"`
	Password     string
	Confirmed    string `gorm:"default:1"`
}
