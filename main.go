package main

import (
	"fillings-back/configs"
	"fillings-back/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	routes.UserRoute(app)

	app.Listen(":3000")
}
