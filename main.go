package main

import (
    "log"
    "package/database"
    "package/routes"
    "github.com/gofiber/fiber/v2")

func main() {
    if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

    api := app.Group("/api")
    routes.UserRoute(api.Group("/users"))

    app.Listen(":3000")
}