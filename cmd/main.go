// main.go
package main

import (
	"github.com/divrhino/divrhino-trivia/database"
	"github.com/divrhino/divrhino-trivia/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Initialisation de la connexion à la base de données
	database.ConnectDb()

	// Configuration des routes
	routes.SetupUserRoutes(app)
	routes.SetupInteractionRoutes(app)
	routes.SetupFollowRoutes(app)
	routes.SetupFavoriteRoutes(app)

	// Démarrage du serveur
	app.Listen(":3000")
}
