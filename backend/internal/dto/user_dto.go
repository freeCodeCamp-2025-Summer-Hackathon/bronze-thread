package dto

import "time"

// UserProfileResponse is the response for a public user profile.
type UserProfileResponse struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	AvatarURL    *string   `json:"avatar_url"`
	Bio          *string   `json:"bio"`
	City         *string   `json:"city"`
	State        *string   `json:"state"`
	RatingAvg    float64   `json:"rating_avg"`
	TotalRatings int       `json:"total_ratings"`
	TotalSwaps   int       `json:"total_swaps"`
	CreatedAt    time.Time `json:"created_at"`
}

// UpdateProfileRequest represents the data for updating a user's profile.
type UpdateProfileRequest struct {
	Name           *string  `json:"name" binding:"omitempty,min=2,max=100"`
	Bio            *string  `json:"bio" binding:"omitempty,max=500"`
	Phone          *string  `json:"phone" binding:"omitempty,max=20"`
	AvatarURL      *string  `json:"avatar_url" binding:"omitempty,url"`
	City           *string  `json:"city" binding:"omitempty,max=100"`
	State          *string  `json:"state" binding:"omitempty,max=100"`
	LocationLat    *float64 `json:"location_lat" binding:"omitempty,min=-90,max=90"`
	LocationLng    *float64 `json:"location_lng" binding:"omitempty,min=-180,max=180"`
	SearchRadiusKM *int     `json:"search_radius_km" binding:"omitempty,min=1,max=500"`
}

// UpdateProfileResponse is the response after updating a profile.
type UpdateProfileResponse struct {
	Message string `json:"message"`
}
