package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Name         string         `gorm:"type:varchar(20); not null"`
	MobileNumber string         `gorm:"uniqueIndex;size:15; not null"`
	Email        string         `gorm:"uniqueIndex;size:30; not null"`
	Password     string         `json:"-"`
	Confirmed    *string        `gorm:"default:1"`
}
