package dto

import "time"

// ItemCreateRequest represents the data for creating a new item.
type ItemCreateRequest struct {
	CategoryID     uint     `json:"category_id" binding:"required"`
	Name           string   `json:"name" binding:"required,min=2,max=200"`
	Description    string   `json:"description" binding:"required,min=10,max=2000"`
	Condition      string   `json:"condition" binding:"required,oneof=new like_new good fair poor"`
	EstimatedValue float64  `json:"estimated_value" binding:"required,min=0"`
	Currency       string   `json:"currency" binding:"omitempty,len=3"`
	ImageURLs      []string `json:"image_urls" binding:"required,min=1,max=10"`
}

// ItemCreateResponse is the response after creating an item.
type ItemCreateResponse struct {
	Message string `json:"message"`
	ItemID  uint   `json:"item_id"`
}

// ItemResponse defines the structure for a single item.
type ItemResponse struct {
	ID             uint             `json:"id"`
	Name           string           `json:"name"`
	Description    string           `json:"description"`
	Category       CategoryResponse `json:"category"`
	Condition      string           `json:"condition"`
	EstimatedValue float64          `json:"estimated_value"`
	Currency       string           `json:"currency"`
	ImageURLs      []string         `json:"image_urls"`
	Owner          UserResponse     `json:"owner"`
	Status         string           `json:"status"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
}

// ItemsResponse defines the structure for a list of items.
type ItemsResponse struct {
	Items        []ItemResponse `json:"items"`
	TotalResults int            `json:"total_results"`
}

// ItemUpdateRequest represents the data for updating an existing item.
type ItemUpdateRequest struct {
	Name           *string  `json:"name" binding:"omitempty,min=2,max=200"`
	Description    *string  `json:"description" binding:"omitempty,min=10,max=2000"`
	Condition      *string  `json:"condition" binding:"omitempty,oneof=new like_new good fair poor"`
	EstimatedValue *float64 `json:"estimated_value" binding:"omitempty,min=0"`
	Status         *string  `json:"status" binding:"omitempty,oneof=available pending swapped"`
}

// ItemUpdateResponse is the response after updating an item.
type ItemUpdateResponse struct {
	Message string `json:"message"`
	ItemID  uint   `json:"item_id"`
}

// ItemDeleteResponse is the response after deleting an item.
type ItemDeleteResponse struct {
	Message string `json:"message"`
}
