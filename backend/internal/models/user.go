package models

import (
	"gorm.io/gorm"
	"time"
)

// User represents a user in the system
type User struct {
	ID             uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Email          string         `json:"email" gorm:"unique;not null" binding:"required,email"`
	PasswordHash   string         `json:"-" gorm:"not null"` // Hidden from JSON
	Name           string         `json:"name" gorm:"not null" binding:"required"`
	Phone          *string        `json:"phone" gorm:"null"`
	AvatarURL      *string        `json:"avatar_url" gorm:"null"`
	Bio            *string        `json:"bio" gorm:"null"`
	LocationLat    *float64       `json:"location_lat" gorm:"type:decimal(10,8);null"`
	LocationLng    *float64       `json:"location_lng" gorm:"type:decimal(11,8);null"`
	City           *string        `json:"city" gorm:"null"`
	State          *string        `json:"state" gorm:"null"`
	SearchRadiusKM int            `json:"search_radius_km" gorm:"default:25"`
	RatingAvg      float64        `json:"rating_avg" gorm:"type:decimal(3,2);default:0.0"`
	TotalRatings   int            `json:"total_ratings" gorm:"default:0"`
	TotalSwaps     int            `json:"total_swaps" gorm:"default:0"`
	CreatedAt      time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"` // Soft delete

	// Relationships
	OwnedItems       []Item        `json:"-" gorm:"foreignKey:OwnerID"`
	SwapRequests     []SwapRequest `json:"-" gorm:"foreignKey:RequesterID"`
	ReceivedRequests []SwapRequest `json:"-" gorm:"foreignKey:OwnerID"`
	SentMessages     []ChatMessage `json:"-" gorm:"foreignKey:SenderID"`
	GivenRatings     []UserRating  `json:"-" gorm:"foreignKey:RaterID"`
	ReceivedRatings  []UserRating  `json:"-" gorm:"foreignKey:RatedUserID"`
	Sessions         []UserSession `json:"-" gorm:"foreignKey:UserID"`
}
