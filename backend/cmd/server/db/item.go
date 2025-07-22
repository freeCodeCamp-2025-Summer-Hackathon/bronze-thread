package db

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	OwnerID uint
	CategoryID uint
	Name string
	Description string 

	Owner User
}