package controllers

import (
	"net/http"
	"strconv"

	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/database"
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/dto"
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/models"
	"github.com/gin-gonic/gin"
)

// GetUserProfile handles fetching a public user profile.
func GetUserProfile(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid user ID"})
		return
	}

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "User not found"})
		return
	}

	c.JSON(http.StatusOK, dto.UserProfileResponse{
		ID:           user.ID,
		Name:         user.Name,
		AvatarURL:    user.AvatarURL,
		Bio:          user.Bio,
		City:         user.City,
		State:        user.State,
		RatingAvg:    user.RatingAvg,
		TotalRatings: user.TotalRatings,
		TotalSwaps:   user.TotalSwaps,
		CreatedAt:    user.CreatedAt,
	})
}

// UpdateUserProfile handles updating the authenticated user's profile.
func UpdateUserProfile(c *gin.Context) {
	// Get user ID from the JWT token (set by the IsAuthorized middleware)
	userID, exists := c.Get("sub")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Unauthorized"})
		return
	}

	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Validation failed", Message: err.Error()})
		return
	}

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "User not found"})
		return
	}

	// Update fields if they are provided in the request
	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Bio != nil {
		user.Bio = req.Bio
	}
	// ... (add other fields from UpdateProfileRequest here)

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, dto.UpdateProfileResponse{Message: "Profile updated successfully"})
}
