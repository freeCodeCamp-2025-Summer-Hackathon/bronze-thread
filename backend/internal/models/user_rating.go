package models

import "time"

// UserRating represents a rating between users
type UserRating struct {
	ID                uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	SwapTransactionID uint      `json:"swap_transaction_id" gorm:"not null;index"`
	RaterID           uint      `json:"rater_id" gorm:"not null;index"`
	RatedUserID       uint      `json:"rated_user_id" gorm:"not null;index"`
	Rating            int       `json:"rating" gorm:"not null;check:rating >= 1 AND rating <= 5" binding:"required,min=1,max=5"`
	Review            *string   `json:"review" gorm:"null"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`

	// Relationships
	SwapTransaction SwapTransaction `json:"-" gorm:"foreignKey:SwapTransactionID"`
	Rater           User            `json:"rater" gorm:"foreignKey:RaterID"`
	RatedUser       User            `json:"rated_user" gorm:"foreignKey:RatedUserID"`
}
