package controllers

import (
	"net/http"

	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/cmd/server/services"
	"github.com/gin-gonic/gin"
)

func GetUserDetails(c *gin.Context) {
	id := c.Param("id")

	user, err := services.User.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"user":    user,
	})
}

// type UpdateRequest struct {
// 	Email string `json:"email"`
// }

// func UpdateUserDetails(c *gin.Context) {
// 	var req UpdateRequest
// 	id := c.Param("id")

// 	userInterface, exists := c.Get("currentUser")
// 	if exists == false {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"message": "User not signed in.",
// 		})
// 		return
// 	}

// 	user, ok := userInterface.(db.User)
// 	if !ok {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user type"})
// 		return
// 	}

// 	if strconv.FormatUint(uint64(user.ID), 10) != id {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"message": "You cannot edit other user's profile.",
// 		})
// 		return
// 	}

// 	updatedUser, _ := services.User.UpdateById(user, req)

// 	c.JSON(200, gin.H{
// 		"user": updatedUser,
// 	})
// }
