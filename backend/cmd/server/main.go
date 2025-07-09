package main

import (
	"fmt"
	"log"
	"os"

	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/cmd/server/db"
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/cmd/server/routes"
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

	err = router.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Use the created routes in `routes` directory
	routes.RegisterHealthCheckRoutes(router)
	routes.RegisterAuthenticationRoutes(router)

	// Start server on port 8080
	log.Println("Server running on port http://localhost:" + port)
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
