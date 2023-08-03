package main

import (
	"github.com/Izzy499/crud_api/initializers"
	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	app := fiber.New()

	app.Listen(":5000")
}
