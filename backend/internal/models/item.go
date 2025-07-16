package models

import "time"

// Item represents an item available for swapping
type Item struct {
	ID             uint          `json:"id" gorm:"primaryKey;autoIncrement"`
	OwnerID        uint          `json:"owner_id" gorm:"not null;index"`
	CategoryID     uint          `json:"category_id" gorm:"not null;index"`
	Name           string        `json:"name" gorm:"not null" binding:"required"`
	Description    string        `json:"description" gorm:"not null" binding:"required"`
	Condition      ItemCondition `json:"condition" gorm:"type:varchar(20);not null" binding:"required"`
	EstimatedValue float64       `json:"estimated_value" gorm:"type:decimal(10,2)"`
	Currency       string        `json:"currency" gorm:"type:varchar(3);default:'INR'"`
	ImageURLs      StringArray   `json:"image_urls" gorm:"type:json"`
	Status         ItemStatus    `json:"status" gorm:"type:varchar(20);default:'available';index"`
	CreatedAt      time.Time     `json:"created_at" gorm:"autoCreateTime;index"`
	UpdatedAt      time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	SwappedAt      *time.Time    `json:"swapped_at" gorm:"null"`

	// Relationships
	Owner            User                  `json:"owner" gorm:"foreignKey:OwnerID"`
	Category         Category              `json:"category" gorm:"foreignKey:CategoryID"`
	SwapRequestItems []SwapRequestItem     `json:"-" gorm:"foreignKey:ItemID"`
	RequestedSwaps   []SwapRequest         `json:"-" gorm:"foreignKey:RequestedItemID"`
	ChatRooms        []ChatRoom            `json:"-" gorm:"foreignKey:ItemID"`
	TransactionItems []SwapTransactionItem `json:"-" gorm:"foreignKey:ItemID"`
}
