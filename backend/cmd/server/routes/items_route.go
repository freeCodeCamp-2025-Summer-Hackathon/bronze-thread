package routes

import (
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/cmd/server/controllers"
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/cmd/server/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterItemRoutes(router *gin.Engine) {
	items := router.Group("/items")
	{
		items.GET("/nearby", middlewares.IsAuthorized(), controllers.GetNearByItems)
	}
}