package routes

import (
	"fillings-back/controllers" //add this

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/user", controllers.CreateUser) //add this
}
