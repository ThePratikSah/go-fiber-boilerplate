package main

import (
	"github.com/ThePratikSah/silent-flag/config"
	"github.com/ThePratikSah/silent-flag/database"
	"github.com/ThePratikSah/silent-flag/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// main is the entry point of the Silent Flag application.
//
// It establishes a connection to the database, sets up the Fiber app,
// configures API routes, and starts the server.
//
// No parameters.
// No return values.
func main() {
	database.ConnectDb()
	version := config.Config("API_VERSION", "1")
	prefork := config.Config("FORK", "false") == "true"
	app := fiber.New(fiber.Config{
		Prefork:       prefork,
		CaseSensitive: true,
		AppName:       "Silent Flag",
	})
	app.Use(logger.New())
	app.Use(cors.New())

	api := app.Group("/api")
	routes := api.Group("/v" + version)
	router.UserRoutes(routes)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	app.Listen(":3000")
}
