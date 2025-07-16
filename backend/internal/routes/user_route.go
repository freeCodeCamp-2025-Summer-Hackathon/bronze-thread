package routes

import (
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/controllers"
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.GET("/getUserDetails", middlewares.IsAuthorized(), controllers.GetUserDetails)
	}
}
