package dto

// ItemQueryParams defines the query parameters for filtering items.
type ItemQueryParams struct {
	Latitude   *float64 `form:"latitude" binding:"omitempty,min=-90,max=90"`
	Longitude  *float64 `form:"longitude" binding:"omitempty,min=-180,max=180"`
	Radius     *int     `form:"radius" binding:"omitempty,min=1,max=500"`
	CategoryID *uint    `form:"category_id" binding:"omitempty,min=1"`
	Search     *string  `form:"search" binding:"omitempty,max=100"`
	Condition  *string  `form:"condition" binding:"omitempty"`
	Status     *string  `form:"status" binding:"omitempty,oneof=available,pending,swapped"`
	Page       int      `form:"page" binding:"omitempty,min=1"`
	Limit      int      `form:"limit" binding:"omitempty,min=1,max=100"`
}

// SetDefaults sets default values for ItemQueryParams.
func (params *ItemQueryParams) SetDefaults() {
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.Limit <= 0 {
		params.Limit = 20
	}
}

// SwapRequestQueryParams defines the query parameters for filtering swap requests.
type SwapRequestQueryParams struct {
	Type   *string `form:"type" binding:"omitempty,oneof=sent,received"`
	Status *string `form:"status" binding:"omitempty"`
	Page   int     `form:"page" binding:"omitempty,min=1"`
	Limit  int     `form:"limit" binding:"omitempty,min=1,max=100"`
}

// SetDefaults sets default values for SwapRequestQueryParams.
func (params *SwapRequestQueryParams) SetDefaults() {
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.Limit <= 0 {
		params.Limit = 20
	}
}

// ChatMessagesQueryParams defines the query parameters for fetching chat messages.
type ChatMessagesQueryParams struct {
	Page  int `form:"page" binding:"omitempty,min=1"`
	Limit int `form:"limit" binding:"omitempty,min=1,max=100"`
}

// SetDefaults sets default values for ChatMessagesQueryParams.
func (params *ChatMessagesQueryParams) SetDefaults() {
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.Limit <= 0 {
		params.Limit = 50
	}
}

// RatingsQueryParams defines the query parameters for fetching ratings.
type RatingsQueryParams struct {
	Page  int `form:"page" binding:"omitempty,min=1"`
	Limit int `form:"limit" binding:"omitempty,min=1,max=100"`
}

// SetDefaults sets default values for RatingsQueryParams.
func (params *RatingsQueryParams) SetDefaults() {
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.Limit <= 0 {
		params.Limit = 20
	}
}

// PaginationMeta defines the structure for pagination metadata.
type PaginationMeta struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"total_pages"`
	Total      int `json:"total"`
}
