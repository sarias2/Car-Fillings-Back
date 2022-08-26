package main

import (
	"fillings-back/configs"
	"fillings-back/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// app.Get("/api", func(c *fiber.Ctx) error {
	// 	return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	// })

	configs.ConnectDB()

	routes.UserRoute(app)

	app.Listen(":3000")
}
