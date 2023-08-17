package main

import (
	"github.com/Izzy499/crud_api/initializers"
	routes "github.com/Izzy499/crud_api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.SessionStorage()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept, cookies",
	}))

	routes.Setup(app)

	app.Listen(":5000")
}
