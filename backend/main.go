package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

func main() {
	fmt.Println("Starting server...")

	// Create new Fiber instance
	app := fiber.New(fiber.Config{
		// Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Fiber",
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
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Nyumat!")
	})

	Todos := [] Todo {
		{ID: 1, Title: "Wash Nyumat's Clothes", Done: false, Body: "Body 1"},
		{ID: 2, Title: "Clean Nyumat's Codebase", Done: false, Body: "Body 2"},
		{ID: 3, Title: "Find Nyumat better food", Done: false, Body: "Body 3"},
	}

	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"data":    Todos,
		})
	})

	app.Get("/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("Todo " + id)
	})

	app.Post("/todos", func(c *fiber.Ctx) error {
		type Request struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		}

		var body Request

		if err := c.BodyParser(&body); err != nil {
			return err
		}

		Todos = append(Todos, Todo{
			ID:    len(Todos) + 1,
			Title: body.Title,
			Done:  false,
			Body:  body.Body,
		})

		return c.JSON(fiber.Map{
			"success": true,
			"data":    Todos,
		})
	})


	// Start server with port from .env
	port := func() string {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		return os.Getenv("PORT")
	}

	log.Fatal(app.Listen(":" + port()))
}