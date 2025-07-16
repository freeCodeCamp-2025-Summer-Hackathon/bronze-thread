package routes

import (
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/controllers"
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterItemRoutes(router *gin.Engine) {
	// Group for item-related routes
	itemGroup := router.Group("/items")
	{
		// Public routes, no authentication needed
		itemGroup.GET("/", controllers.GetAllItems)
		itemGroup.GET("/:id", controllers.GetItemByID)

		// Private routes, authentication required
		itemGroup.POST("/", middlewares.IsAuthorized(), controllers.CreateItem)
		itemGroup.PUT("/:id", middlewares.IsAuthorized(), controllers.UpdateItem)
		itemGroup.DELETE("/:id", middlewares.IsAuthorized(), controllers.DeleteItem)
	}
}
