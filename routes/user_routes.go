package routes

import (
	"fillings-back/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {

	app.Post("/user", controllers.CreateUser)
	app.Get("/user/:userId", controllers.GetAUser)
	app.Get("/users", controllers.GetAllUsers)
	app.Put("/user/:userId", controllers.EditAUser)
	app.Delete("/user/:userId", controllers.DeleteAUser)
}
