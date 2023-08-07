package routes

import (
	"github.com/Izzy499/crud_api/controllers"
	"github.com/Izzy499/crud_api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/login", controllers.Login)
	api.Post("/register", controllers.Register)
	api.Post("/verify", controllers.VerifyUser)

	auth := api.Group("/auth", middlewares.LoggedIn())
	auth.Get("/me", controllers.GetUserDetails)
	auth.Get("/user/:id", controllers.GetUserById)
}
