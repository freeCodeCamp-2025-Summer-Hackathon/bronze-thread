package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/cmd/server/db"
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/cmd/server/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	err = db.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	} else {
		fmt.Println("Connected to Database.")
	}

	// Create Gin router
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	err = router.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Use the created routes in `routes` directory
	routes.RegisterHealthCheckRoutes(router)
	routes.RegisterAuthenticationRoutes(router)
	routes.RegisterUserRoutes(router)

	// Start server on port 8080
	log.Println("Server running on port http://localhost:" + port)
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
