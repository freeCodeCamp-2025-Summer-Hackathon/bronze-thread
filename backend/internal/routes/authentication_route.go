package routes

import (
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAuthenticationRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		// Endpoint for new user registration
		authGroup.POST("/register", controllers.Register)

		// Endpoint for user sign-in
		authGroup.POST("/signin", controllers.Signin)
	}
}
