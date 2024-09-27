package router

import (
	"github.com/ThePratikSah/silent-flag/controller"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Get("/users", controller.GetUsers)
	app.Get("/users/:id", controller.GetUser)
	app.Post("/users", controller.CreateUser)
	app.Put("/users/:id", controller.UpdateUser)
	app.Delete("/users/:id", controller.DeleteUser)
}
