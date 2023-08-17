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

	auth := api.Group("/auth", middlewares.LoggedIn())
	auth.Get("/generate-token", controllers.GenerateVerificationToken)
	auth.Get("/post", controllers.GetAllPost)
	auth.Get("/me", controllers.GetUserDetails)
	auth.Get("/user/:id", controllers.GetUserById)
	auth.Post("/verify", controllers.VerifyUser)
	auth.Post("/post", controllers.CreatePost)
	auth.Post("/comment", controllers.AddComment)
	auth.Get("/logout", controllers.Logout)

}
