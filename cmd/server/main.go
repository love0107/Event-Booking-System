package main

import (
	"event-booking-backend/internal/handlers"
	"event-booking-backend/pkg/db"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize the database
	db.InitDB()

	// Create a Gin router
	router := gin.Default()

	// Define routes
	router.GET("/events", handlers.GetEvents)
	router.POST("/events", handlers.CreateEvent)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}
	log.Println("Server running on port:", port)
	router.Run(":" + port)
}
