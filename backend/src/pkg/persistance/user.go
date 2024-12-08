package persistance

import (
	"event-booking-backend/src/internal/models"

	"github.com/labstack/gommon/log"
)

func CreateUser(user *models.User) (int64, error) {
	// Create user
	log.Printf("Creating user: %v", user)
	id, err := user.CreateUser()
	if err != nil {
		return 0, err
	}
	return id, nil
}
