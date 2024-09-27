package main

import (
	"github.com/ThePratikSah/silent-flag/database"
	"github.com/ThePratikSah/silent-flag/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.ConnectDb()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	router.UserRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	app.Listen(":3000")
}
