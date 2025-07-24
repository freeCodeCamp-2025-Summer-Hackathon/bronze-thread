package db

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string `gorm:"uniqueIndex"`
	Password string
	Phone string
	AvatarURL string
	Bio string
	
	// location related fields
	LocationLat float64 `gorm:"type:decimal(10,8)"`
	LocationLng float64 `gorm:"type:decimal(11,8)"`
	City string
	State string
	SearchRadiusKm int

	// fields related to items and rating
	RatingAvg  float64 `gorm:"type:decimal(3,2)"`
	TotalRatings int
	TotalSwaps int
}

type Item struct {
	gorm.Model
	OwnerID uint
	CategoryID uint
	Name string
	Description string 

	Owner User
}