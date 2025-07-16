package routes

import (
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/controllers"
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	// Group for user-related routes
	userGroup := router.Group("/users")
	{
		// Get a public user profile by ID - No authentication required
		// Corresponds to GET /users/{user_id} in our API design
		userGroup.GET("/:id", controllers.GetUserProfile)

		// A nested group for the authenticated user's own profile
		profileGroup := userGroup.Group("/profile")

		profileGroup.Use(middlewares.IsAuthorized()) // Protect this route
		{
			// Corresponds to PUT /users/profile in your API design
			profileGroup.PUT("/", controllers.UpdateUserProfile)
		}
	}
}
