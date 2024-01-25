// interaction_routes.go
package routes

import (
	"github.com/divrhino/divrhino-trivia/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupInteractionRoutes(app *fiber.App) {
	// Routes publiques
	app.Post("/articles/:articleid/like", handlers.LikeArticle)
	app.Post("/articles/:articleid/dislike", handlers.DislikeArticle)

	// Routes nécessitant une connexion
	authorized := app.Group("/user", handlers.RequireAuthMiddleware)

	authorized.Get("/articles", handlers.GetAllArticles)
	authorized.Post("/articles", handlers.CreateArticle)
	authorized.Post("/articles/:articleid/comment", handlers.AddComment)
	authorized.Delete("/articles/:articleid", handlers.DeleteArticle)
}
