package routes

import (
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/cmd/server/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterHealthCheckRoutes(router *gin.Engine) {
	userGroup := router.Group("/health-check")
	{
		userGroup.GET("/hello-world", controllers.HelloWorld)
	}
}
