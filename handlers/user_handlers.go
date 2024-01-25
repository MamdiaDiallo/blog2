package handlers

import (
	"github.com/MamdiaDiallo/blog2/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/MamdiaDiallo/blog2/models"
)

// GetUser retrieves a user by ID
func GetUser(c *fiber.Ctx) error {
	userId := c.Params("id")

	var user models.User
	result := database.DB.Db.First(&user, userId)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

// CreateUser creates a new user
func CreateUser(c *fiber.Ctx) error {
	var newUser models.User

	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	result := database.DB.Db.Create(&newUser)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.JSON(newUser)
}
