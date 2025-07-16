package models

import "time"

// Category represents item categories
type Category struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"unique;not null" binding:"required"`
	Description *string   `json:"description" gorm:"null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`

	// Relationships
	Items []Item `json:"-" gorm:"foreignKey:CategoryID"`
}
