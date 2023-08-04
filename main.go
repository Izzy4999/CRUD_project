package main

import (
	"github.com/Izzy499/crud_api/initializers"
	routes "github.com/Izzy499/crud_api/routes"
	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.SessionStorage()
	app := fiber.New()

	routes.Setup(app)

	app.Listen(":5000")
}
