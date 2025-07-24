package controllers

import (
	"net/http"
	"strconv"

	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/cmd/server/db"
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/cmd/server/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GetNearByItems(c* gin.Context) {
	userData, exists := c.Get("user")
    if !exists {
        c.JSON(401, gin.H{"error": "unauthorized"})
        return
    }
	claims, ok := userData.(jwt.MapClaims)
	if !ok {
		c.JSON(401, gin.H{"error": "invalid token"})
		return
	}
	email := claims["email"].(string)
	var user db.User
	result := db.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User Not found"})
		return
	}

	latitudeString := c.Query("lat")
	longitudeString := c.Query("lng")
	latitude, latitudeError := strconv.ParseFloat(latitudeString, 64)
	longitude, longitudeError := strconv.ParseFloat(longitudeString, 64)

	if latitudeError != nil {
		latitude = user.LocationLat
	}

	if longitudeError != nil {
		longitude = user.LocationLng
	}

	users, err := services.GetNearByUsers(latitude, longitude, float64(user.SearchRadiusKm))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch nearby users"})
		return 
	}

	userItemsMapping, err := services.GetUserItemsMapping(users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{	
		"items": userItemsMapping,
	})
}