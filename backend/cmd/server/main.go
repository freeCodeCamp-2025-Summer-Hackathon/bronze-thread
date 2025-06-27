package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	// Create Gin router
	router := gin.Default()

	router.SetTrustedProxies(nil)

	// Define the GET /hello endpoint
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	// Start server on port 8080
	log.Println("Server running on port http://localhost:" + port)
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
