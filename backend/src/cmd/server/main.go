package main

import (
	"event-booking-backend/src/pkg/auth"
	"event-booking-backend/src/pkg/db"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// This is where the program starts
	// Log the current working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}
	log.Printf("Current working directory: %s", wd)

	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	// Initialize the database
	db.InitDB()

	// Create a Gin router
	router := gin.Default()

	// Register the signup and login routes
	router.POST("/signup", auth.SignUp)
	//  router.POST("/login", login)

	router.Run(":8080")

}
