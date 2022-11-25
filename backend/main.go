package main

import (
	"backend/configs"
	"backend/routes"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Create new Fiber instance
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "FiberM8",
		AppName:       "~ GolangGoBackend ~",
	})

	// Middleware 
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept",
		},
	))

	// Routes
	routes.UserRoutes(app)

	// MongoDB Connection
	configs.MongoConnect()

	fmt.Printf("%sSuccess: Hi Nyumat, Welcome back to MongoDB!%s ", configs.ColorGreen, configs.ColorReset)

	// Listen to port from .env
	port := configs.PORT()
	log.Fatal(app.Listen(":" + port))
}