package repository

import (
	"github.com/MamdiaDiallo/blog2/database"
	"github.com/gofiber/fiber/v2"
	"github.com/MamdiaDiallo/blog2/handlers"
)

// GetUserByID retrieves a user by ID
func GetUserByID(userID int) (models.User, error) {
	var user models.User
	result := database.DB.Db.First(&user, userID)
	return user, result.Error
}

// CreateUser creates a new user
func CreateUser(newUser *models.User) error {
	result := database.DB.Db.Create(newUser)
	return result.Error
}
