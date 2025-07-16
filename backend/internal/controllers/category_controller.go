package controllers

import (
	"net/http"

	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/database"
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/dto"
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/models"
	"github.com/gin-gonic/gin"
)

// GetAllCategories handles the retrieval of all item categories.
func GetAllCategories(c *gin.Context) {
	var categories []models.Category
	if err := database.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to retrieve categories"})
		return
	}

	var categoryResponses []dto.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, dto.CategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	c.JSON(http.StatusOK, dto.CategoriesResponse{Categories: categoryResponses})
}
