package dto

import "time"

// ChatRoomResponse defines the structure for a single chat room.
type ChatRoomResponse struct {
	ID            uint          `json:"id"`
	OtherUser     UserResponse  `json:"other_user"`
	Item          *ItemResponse `json:"item"`
	SwapRequestID *uint         `json:"swap_request_id"`
	LastMessage   *LastMessage  `json:"last_message"`
	IsActive      bool          `json:"is_active"`
	LastMessageAt time.Time     `json:"last_message_at"`
}

// LastMessage defines the structure for the last message in a chat room.
type LastMessage struct {
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

// ChatMessageResponse defines the structure for a single chat message.
type ChatMessageResponse struct {
	ID        uint         `json:"id"`
	Sender    UserResponse `json:"sender"`
	Message   string       `json:"message"`
	CreatedAt time.Time    `json:"created_at"`
}

// ChatMessagesResponse defines the structure for a paginated list of chat messages.
type ChatMessagesResponse struct {
	Messages   []ChatMessageResponse `json:"messages"`
	Page       int                   `json:"page"`
	Limit      int                   `json:"limit"`
	TotalPages int                   `json:"total_pages"`
}

// ChatRoomsResponse defines the structure for a list of chat rooms.
type ChatRoomsResponse struct {
	ChatRooms []ChatRoomResponse `json:"chat_rooms"`
}

// SendMessageRequest represents the data for sending a new message.
type SendMessageRequest struct {
	Message string `json:"message" binding:"required,min=1,max=1000"`
}

// SendMessageResponse is the response after sending a message.
type SendMessageResponse struct {
	Message   string `json:"message"`
	MessageID uint   `json:"message_id"`
}

// CreateChatRoomRequest represents the data for creating a new chat room.
type CreateChatRoomRequest struct {
	OtherUserID uint `json:"other_user_id" binding:"required"`
	ItemID      uint `json:"item_id" binding:"required"`
}

// CreateChatRoomResponse is the response after creating a chat room.
type CreateChatRoomResponse struct {
	Message    string `json:"message"`
	ChatRoomID uint   `json:"chat_room_id"`
}
