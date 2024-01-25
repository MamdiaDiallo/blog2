// routes.go
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/divrhino/divrhino-trivia/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/login", handlers.Login)
	app.Post("/user", handlers.CreateUser)
	app.Get("/home", handlers.GetHome)
}
