package server

import (
	"log"
	"os"

	"mywebapi/proceed"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// Start initializes the Gin server with the specified settings
func Start(db *gorm.DB) {
	// Load environment variables from .env file
	err := godotenv.Load("./info/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Retrieve server settings from environment variables
	port := os.Getenv("SERVER_PORT")
	mode := os.Getenv("SERVER_MODE")

	// Set the Gin mode (debug, release, test)
	gin.SetMode(mode)

	// Initialize Gin router
	router := gin.Default()

	// Set up routes and pass the database connection
	proceed.SetupRoutes(router, db)

	// Start the server
	if port == "" {
		port = "8080" // Default to port 8080 if not set
	}
	err = router.Run(":" + port)
	if err != nil {
		log.Fatalf("err starting server : %v", err)
	}
	log.Fatalf("server listening on port : %s", port)
}
