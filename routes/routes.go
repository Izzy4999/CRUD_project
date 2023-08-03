package routes

import (
	"github.com/Izzy499/crud_api/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/login", nil)
	api.Post("/register", controllers.Register)
}
