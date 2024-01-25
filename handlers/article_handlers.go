// article_handlers.go
package handlers

import (
	"github.com/divrhino/divrhino-trivia/database"
	"github.com/divrhino/divrhino-trivia/models"
	"github.com/gofiber/fiber/v2"
)

// GetAllArticles récupère tous les articles d'un blogueur
func GetAllArticles(c *fiber.Ctx) error {
	// Vérifier si l'utilisateur est connecté
	user := getCurrentUser(c)
	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Récupérer tous les articles de l'utilisateur
	articles := []models.Article{}
	database.DB.Db.Where("user_id = ?", user.ID).Find(&articles)

	return c.Status(fiber.StatusOK).JSON(articles)
}

// CreateArticle ajoute un article sur le blog de l'utilisateur
func CreateArticle(c *fiber.Ctx) error {
	// Vérifier si l'utilisateur est connecté
	user := getCurrentUser(c)
	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Parsage du JSON dans le corps de la requête
	article := new(models.Article)
	if err := c.BodyParser(article); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Ajout de l'ID de l'utilisateur à l'article
	article.UserID = user.ID

	// Enregistrement de l'article dans la base de données
	database.DB.Db.Create(&article)

	return c.Status(fiber.StatusCreated).JSON(article)
}

// AddComment ajoute un commentaire à un article
func AddComment(c *fiber.Ctx) error {
	// Vérifier si l'utilisateur est connecté
	user := getCurrentUser(c)
	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Parsage du JSON dans le corps de la requête
	comment := new(models.Comment)
	if err := c.BodyParser(comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Ajout de l'ID de l'utilisateur à l'article
	comment.UserID = user.ID

	// Enregistrement du commentaire dans la base de données
	database.DB.Db.Create(&comment)

	return c.Status(fiber.StatusCreated).JSON(comment)
}

// LikeArticle permet à l'utilisateur de liker un article
func LikeArticle(c *fiber.Ctx) error {
	// Vérifier si l'utilisateur est connecté
	user := getCurrentUser(c)
	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Récupérer l'ID de l'article depuis les paramètres de la requête
	articleID := c.Params("articleid")

	// Mettre à jour le nombre de likes de l'article dans la base de données
	database.DB.Db.Model(&models.Article{}).Where("id = ?", articleID).Update("likes", gorm.Expr("likes + ?", 1))

	return c.SendStatus(fiber.StatusOK)
}

// DislikeArticle permet à l'utilisateur de disliker un article
func DislikeArticle(c *fiber.Ctx) error {
	// Vérifier si l'utilisateur est connecté
	user := getCurrentUser(c)
	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Récupérer l'ID de l'article depuis les paramètres de la requête
	articleID := c.Params("articleid")

	// Mettre à jour le nombre de likes de l'article dans la base de données
	database.DB.Db.Model(&models.Article{}).Where("id = ?", articleID).Update("likes", gorm.Expr("likes - ?", 1))

	return c.SendStatus(fiber.StatusOK)
}

// DeleteArticle permet à l'utilisateur de supprimer un article
func DeleteArticle(c *fiber.Ctx) error {
	// Vérifier si l'utilisateur est connecté
	user := getCurrentUser(c)
	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Récupérer l'ID de l'article depuis les paramètres de la requête
	articleID := c.Params("articleid")

	// Supprimer l'article de la base de données
	database.DB.Db.Delete(&models.Article{}, articleID)

	return c.SendStatus(fiber.StatusOK)
}
