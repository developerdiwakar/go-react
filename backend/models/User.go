package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Name         string         `gorm:"type:varchar(20); not null" json:"name"`
	MobileNumber string         `gorm:"uniqueIndex;size:15; not null" json:"mobile_number"`
	Email        string         `gorm:"uniqueIndex;size:30; not null" json:"email"`
	Password     string         `json:"-"`
	Confirmed    *string        `gorm:"default:1" json:"confirmed"`
}
