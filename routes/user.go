package routes

import (
	"package/controllers" // replace

	"github.com/gofiber/fiber/v2"
)

func UserRoute(route fiber.Router) {
	route.Get("/", controllers.GetUsers)
	route.Delete("/:id", controllers.GetUsers)
	route.Post("/", controllers.AddUser)
}
