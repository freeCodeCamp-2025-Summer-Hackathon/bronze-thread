package routes

import (
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAuthenticationRoutes(router *gin.Engine) {
	userGroup := router.Group("/auth")
	{
		userGroup.POST("/signin", controllers.Signin)
		userGroup.POST("/register", controllers.Register)
	}
}
