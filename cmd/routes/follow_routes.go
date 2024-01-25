// follow_routes.go
package routes

import (
	"github.com/divrhino/divrhino-trivia/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupFollowRoutes(app *fiber.App) {
	app.Get("/followers/:userid", handlers.GetFollowers)
	app.Post("/follow/:userId", handlers.FollowUser)
	app.Post("/unfollow/:userId", handlers.UnfollowUser)
}
