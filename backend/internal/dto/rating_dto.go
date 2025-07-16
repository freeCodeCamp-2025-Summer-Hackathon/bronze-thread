package dto

import "time"

// UserRatingRequest represents the data for submitting a new user rating.
type UserRatingRequest struct {
	SwapTransactionID uint    `json:"swap_transaction_id" binding:"required"`
	RatedUserID       uint    `json:"rated_user_id" binding:"required"`
	Rating            int     `json:"rating" binding:"required,min=1,max=5"`
	Review            *string `json:"review" binding:"omitempty,max=500"`
}

// SubmitRatingResponse is the response after submitting a rating.
type SubmitRatingResponse struct {
	Message  string `json:"message"`
	RatingID uint   `json:"rating_id"`
}

// UserRatingResponse defines the structure for a single user rating.
type UserRatingResponse struct {
	ID        uint         `json:"id"`
	Rater     UserResponse `json:"rater"`
	Rating    int          `json:"rating"`
	Review    *string      `json:"review"`
	CreatedAt time.Time    `json:"created_at"`
}

// UserRatingsResponse defines the structure for a paginated list of user ratings.
type UserRatingsResponse struct {
	Ratings       []UserRatingResponse `json:"ratings"`
	AverageRating float64              `json:"average_rating"`
	TotalRatings  int                  `json:"total_ratings"`
	Page          int                  `json:"page"`
	Limit         int                  `json:"limit"`
	TotalPages    int                  `json:"total_pages"`
}
