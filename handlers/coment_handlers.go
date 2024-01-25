// comment_handlers.go
package handlers

import (
	"github.com/divrhino/divrhino-trivia/database"
	"github.com/divrhino/divrhino-trivia/models"
	"github.com/gofiber/fiber/v2"
)

func AddComment(c *fiber.Ctx) error {
	comment := new(models.Comment)
	if err := c.BodyParser(comment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&comment)

	return c.Status(fiber.StatusOK).JSON(comment)
}

