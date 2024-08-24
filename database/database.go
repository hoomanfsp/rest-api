package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type JoinUsRequest struct {
	ID          uint   `gorm:"primaryKey"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Description string `json:"description"`
	Email       string `json:"email" binding:"required,email"`
}

// InitDatabase loads the environment variables, connects to the database, and initializes the tables
func InitDatabase() (*gorm.DB, error) {
	// Load environment variables from .env file
	err := godotenv.Load("./info/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Retrieve database connection info from environment variables
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	// Construct the Data Source Name (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	// Open a connection to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Database connection established")

	// AutoMigrate will create/update the table structure in the database
	db.AutoMigrate(&JoinUsRequest{})

	return db, nil
}
