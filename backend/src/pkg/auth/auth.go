package auth

import (
	"event-booking-backend/src/internal/models"
	"event-booking-backend/src/pkg/persistance"
	"event-booking-backend/src/pkg/utils"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var newUser *models.User
	if err := c.ShouldBindJSON(newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate mobile number
	if !regexp.MustCompile(`^\d{10}$`).MatchString(newUser.Mobile) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mobile number"})
		return
	}

	newUser.Mobile = "91" + newUser.Mobile

	// Hash the password
	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	newUser.Password = hashedPassword

	// Create the new user
	userId, err := persistance.CreateUser(newUser)
	if err != nil || userId == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to create new user",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"userId":     userId,
		"message":    "User created successfully",
		"statusCode": 200,
	})
}
