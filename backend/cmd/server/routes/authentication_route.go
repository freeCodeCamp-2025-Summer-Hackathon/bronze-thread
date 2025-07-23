package routes

import (
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/cmd/server/controllers"
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/cmd/server/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterAuthenticationRoutes(router *gin.Engine) {
	userGroup := router.Group("/auth")
	{
		userGroup.POST("/signin", controllers.Signin)
		userGroup.POST("/register", controllers.Register)
		userGroup.POST("/signout", middlewares.IsAuthorized(), controllers.SignOut)
		userGroup.GET("/me", middlewares.IsAuthorized(), controllers.GetSignedInUser)
	}
}
