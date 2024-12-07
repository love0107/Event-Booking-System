package handlers

import (
	"event-booking-backend/internal/models"
	"event-booking-backend/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetEvents(c *gin.Context) {
	var events []models.Event
	if result := db.DB.Find(&events); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

func CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if result := db.DB.Create(&event); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, event)
}
