package routes

import (
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/controllers"
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterSwapRequestRoutes(router *gin.Engine) {
	// Group for swap request-related routes
	swapRequestGroup := router.Group("/swap-requests")
	// All routes in this group require authentication
	swapRequestGroup.Use(middlewares.IsAuthorized())
	{
		swapRequestGroup.POST("/", controllers.CreateSwapRequest)
		swapRequestGroup.GET("/", controllers.GetSwapRequests)
		swapRequestGroup.PUT("/:id", controllers.UpdateSwapRequest)
	}
}
