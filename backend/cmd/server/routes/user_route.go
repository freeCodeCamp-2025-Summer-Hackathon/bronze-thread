package routes

import (
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/cmd/server/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.GET("/:id", controllers.GetUserDetails)
		// userGroup.PATCH("/:id", middlewares.IsAuthorized(), controllers.UpdateUserDetails)
	}
}
