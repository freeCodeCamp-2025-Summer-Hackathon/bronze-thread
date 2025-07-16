package routes

import (
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(router *gin.Engine) {
	// Group for category-related routes
	categoryGroup := router.Group("/categories")
	{
		// Get all categories - No authentication required
		// Corresponds to GET /categories in your API design
		categoryGroup.GET("/", controllers.GetAllCategories)
	}
}
