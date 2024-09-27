package router

import (
	"github.com/ThePratikSah/silent-flag/controller"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app fiber.Router) {
	app.Route("/users", func(router fiber.Router) {
		router.Get("/", controller.GetUsers)
		router.Get("/:id", controller.GetUser)
		router.Post("/", controller.CreateUser)
		router.Put("/:id", controller.UpdateUser)
		router.Delete("/:id", controller.DeleteUser)
	})
}
