package models

import (
	"gorm.io/gorm"
	"time"
)

// SwapRequest represents a swap request between users
type SwapRequest struct {
	ID              uint              `json:"id" gorm:"primaryKey;autoIncrement"`
	RequesterID     uint              `json:"requester_id" gorm:"not null;index"`
	OwnerID         uint              `json:"owner_id" gorm:"not null;index"`
	RequestedItemID uint              `json:"requested_item_id" gorm:"not null;index"`
	Status          SwapRequestStatus `json:"status" gorm:"type:varchar(20);default:'pending';index"`
	CreatedAt       time.Time         `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time         `json:"updated_at" gorm:"autoUpdateTime"`
	ExpiresAt       time.Time         `json:"expires_at" gorm:"index"`

	// Relationships
	Requester        User              `json:"requester" gorm:"foreignKey:RequesterID"`
	Owner            User              `json:"owner" gorm:"foreignKey:OwnerID"`
	RequestedItem    Item              `json:"requested_item" gorm:"foreignKey:RequestedItemID"`
	SwapRequestItems []SwapRequestItem `json:"offered_items" gorm:"foreignKey:SwapRequestID"`
	ChatRooms        []ChatRoom        `json:"-" gorm:"foreignKey:SwapRequestID"`
	SwapTransaction  *SwapTransaction  `json:"swap_transaction" gorm:"foreignKey:SwapRequestID"`
}

// SwapRequestItem represents items involved in a swap request
type SwapRequestItem struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	SwapRequestID uint      `json:"swap_request_id" gorm:"not null;index"`
	ItemID        uint      `json:"item_id" gorm:"not null;index"`
	ItemType      string    `json:"item_type" gorm:"type:varchar(20);not null"` // "requested" or "offered"
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`

	// Relationships
	SwapRequest SwapRequest `json:"-" gorm:"foreignKey:SwapRequestID"`
	Item        Item        `json:"item" gorm:"foreignKey:ItemID"`
}

// BeforeCreate hook for SwapRequest to set expiry
func (sr *SwapRequest) BeforeCreate(tx *gorm.DB) error {
	if sr.ExpiresAt.IsZero() {
		sr.ExpiresAt = time.Now().Add(7 * 24 * time.Hour) // 7 days from now
	}
	return nil
}

// IsExpired checks if a swap request has expired.
func (sr *SwapRequest) IsExpired() bool {
	return time.Now().After(sr.ExpiresAt)
}
