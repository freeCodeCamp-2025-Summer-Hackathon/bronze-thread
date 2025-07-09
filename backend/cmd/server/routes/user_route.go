package routes

import (
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/cmd/server/controllers"
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/cmd/server/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.GET("/getUserDetails", middlewares.IsAuthorized(), controllers.GetUserDetails)
	}
}
