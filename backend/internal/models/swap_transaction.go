package models

import "time"

// SwapTransaction represents a confirmed swap
type SwapTransaction struct {
	ID            uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	SwapRequestID uint       `json:"swap_request_id" gorm:"not null;unique"`
	User1ID       uint       `json:"user1_id" gorm:"not null;index"`
	User2ID       uint       `json:"user2_id" gorm:"not null;index"`
	CreatedAt     time.Time  `json:"created_at" gorm:"autoCreateTime"`
	CompletedAt   *time.Time `json:"completed_at" gorm:"null"`

	// Relationships
	SwapRequest      SwapRequest           `json:"swap_request" gorm:"foreignKey:SwapRequestID"`
	User1            User                  `json:"user1" gorm:"foreignKey:User1ID"`
	User2            User                  `json:"user2" gorm:"foreignKey:User2ID"`
	TransactionItems []SwapTransactionItem `json:"transaction_items" gorm:"foreignKey:SwapTransactionID"`
	Ratings          []UserRating          `json:"ratings" gorm:"foreignKey:SwapTransactionID"`
}

// SwapTransactionItem represents items in a completed swap
type SwapTransactionItem struct {
	ID                uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	SwapTransactionID uint      `json:"swap_transaction_id" gorm:"not null;index"`
	ItemID            uint      `json:"item_id" gorm:"not null;index"`
	UserID            uint      `json:"user_id" gorm:"not null;index"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`

	// Relationships
	SwapTransaction SwapTransaction `json:"-" gorm:"foreignKey:SwapTransactionID"`
	Item            Item            `json:"item" gorm:"foreignKey:ItemID"`
	User            User            `json:"user" gorm:"foreignKey:UserID"`
}
