package routes

import (
	"github.com/gofiber/fiber/v2"
	"backend/controllers"
)

// UserRoutes registers all user routes
func UserRoutes(app *fiber.App) {
	// Test GET '/' Endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg": "Hello Nyumat!",
		})
	})

	app.Get("/api/welcome", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg": "Welcome to the GolangGoBackend!",
		})
	})

	app.Get("/api/users/:id", controllers.GetUser)
	app.Get("/api/user/:name", controllers.GetUserByName)
	app.Post("/api/users", controllers.CreateUser)
	// app.Get("/api/users", controllers.GetUsers)
	// app.Put("/api/users/:id", controllers.UpdateUser)
	// app.Delete("/api/users/:id", controllers.DeleteUser)
}