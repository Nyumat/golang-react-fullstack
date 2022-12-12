package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

// UserRoutes registers all user routes
func UserRoutes(app *fiber.App) {
	app.Get("/api/users/:id", controllers.GetUser)
	app.Get("/api/user/:name", controllers.GetUserByName)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users", controllers.GetUsers)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)
}