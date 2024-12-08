package main

import (
	"event-booking-backend/src/pkg/db"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// This is where the program starts
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize the database
	db.InitDB()


	// Create a Gin router
	router := gin.Default()

	// Define the routes


}
