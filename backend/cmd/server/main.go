package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Create Gin router
	r := gin.Default()

	// Define the GET /hello endpoint
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	// Start server on port 8080
	r.Run(":8080")
}
