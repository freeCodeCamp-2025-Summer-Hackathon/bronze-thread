package models

import "time"

// ChatRoom represents a chat session between users
type ChatRoom struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	SwapRequestID *uint     `json:"swap_request_id" gorm:"null;index"`
	ItemID        *uint     `json:"item_id" gorm:"null;index"`
	User1ID       uint      `json:"user1_id" gorm:"not null;index"`
	User2ID       uint      `json:"user2_id" gorm:"not null;index"`
	IsActive      bool      `json:"is_active" gorm:"default:true"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	LastMessageAt time.Time `json:"last_message_at" gorm:"default:CURRENT_TIMESTAMP"`

	// Relationships
	SwapRequest *SwapRequest  `json:"swap_request" gorm:"foreignKey:SwapRequestID"`
	Item        *Item         `json:"item" gorm:"foreignKey:ItemID"`
	User1       User          `json:"user1" gorm:"foreignKey:User1ID"`
	User2       User          `json:"user2" gorm:"foreignKey:User2ID"`
	Messages    []ChatMessage `json:"messages" gorm:"foreignKey:ChatRoomID"`
}

// ChatMessage represents a message in a chat room
type ChatMessage struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	ChatRoomID uint      `json:"chat_room_id" gorm:"not null;index"`
	SenderID   uint      `json:"sender_id" gorm:"not null;index"`
	Message    string    `json:"message" gorm:"not null" binding:"required"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime;index"`

	// Relationships
	ChatRoom ChatRoom `json:"-" gorm:"foreignKey:ChatRoomID"`
	Sender   User     `json:"sender" gorm:"foreignKey:SenderID"`
}

// GetOtherUser is a helper method to get the other user in a chat room
func (cr *ChatRoom) GetOtherUser(currentUserID uint) *User {
	if cr.User1ID == currentUserID {
		return &cr.User2
	}
	return &cr.User1
}
