// user_routes.go
package routes

import (
	"github.com/divrhino/divrhino-trivia/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	app.Post("/login", handlers.Login)
	app.Post("/user", handlers.CreateUser)
	app.Get("/home", handlers.GetRecentArticles)
	app.Get("/user/:userid/profile", handlers.GetUserProfile)
	app.Get("/articles", handlers.GetUserArticles)
}
