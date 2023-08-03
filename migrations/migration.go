package main

import (
	"github.com/Izzy499/crud_api/initializers"
	"github.com/Izzy499/crud_api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
