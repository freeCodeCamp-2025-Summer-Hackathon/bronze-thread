package db

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	Username  string
	Email     string         `gorm:"uniqueIndex"`
	Password  string         `json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
