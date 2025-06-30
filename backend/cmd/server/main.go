package main

import (
	"log"
	"os"

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

	// Create Gin router
	router := gin.Default()

	err = router.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Use the created routes in `routes` directory
	routes.RegisterHealthCheckRoutes(router)

	// Start server on port 8080
	log.Println("Server running on port http://localhost:" + port)
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
