package dto

import "time"

// RegisterRequest represents the data required for user registration.
type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// RegisterResponse is the response sent after a successful user registration.
type RegisterResponse struct {
	Message string `json:"message"`
	UserID  uint   `json:"user_id"`
}

// LoginRequest represents the data required for user login.
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse is the response sent after a successful user login.
type LoginResponse struct {
	Message string       `json:"message"`
	Token   string       `json:"token"`
	User    UserResponse `json:"user"`
}

// User DTOs
type UserResponse struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Bio          *string   `json:"bio"`
	AvatarURL    *string   `json:"avatar_url"`
	City         *string   `json:"city"`
	State        *string   `json:"state"`
	RatingAvg    float64   `json:"rating_avg"`
	TotalRatings int       `json:"total_ratings"`
	TotalSwaps   int       `json:"total_swaps"`
	CreatedAt    time.Time `json:"created_at"`
}

// LogoutResponse is the response sent after a successful user logout.
type LogoutResponse struct {
	Message string `json:"message"`
}
