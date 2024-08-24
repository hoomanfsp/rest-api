package main

import (
	"log"
	"mywebapi/database"
	"mywebapi/server"
)

func main() {
	// Initialize the database connection
	db, err := database.InitDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Start the server with the initialized routes
	server.Start(db)
}
