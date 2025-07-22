package controllers

import (
	"net/http"
	"strconv"

	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/cmd/server/services"
	"github.com/gin-gonic/gin"
)

func GetNearByItems(c* gin.Context) {
	latitudeString := c.Query("lat")
	longitudeString := c.Query("lng")

	latitude, latitudeError := strconv.ParseFloat(latitudeString, 64)
	longitude, longitudeError := strconv.ParseFloat(longitudeString, 64)

	if latitudeError != nil || longitudeError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid coordinates"})
		return 
	}

	users, err := services.GetNearByUsers(latitude, longitude, 78.10131)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch nearby users"})
		return 
	}


	c.JSON(http.StatusOK, gin.H{	
		"latitude": latitude,
		"longitude": longitude,
	})
}