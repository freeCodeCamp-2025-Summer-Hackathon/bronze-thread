package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetNearByItems(c* gin.Context) {
	lat := c.Query("lat")
	lng := c.Query("lng")

	c.JSON(http.StatusOK, gin.H{
		"latitude": lat,
		"longitude": lng,
	})
}