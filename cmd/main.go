package main

import (
	"github.com/MamdiaDiallo/blog2"
	"github.com/gofiber/fiber/v2"
	"github.com/MamdiaDiallo/blog2"

)

func setupRoutes(app *fiber.App) {
	app.Get("/user/:id", handlers.GetUser)
	app.Post("/user", handlers.CreateUser)
}

func main() {
	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
