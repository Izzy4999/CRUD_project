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
	initializers.DB.AutoMigrate(&models.Verify_Email{})
	initializers.DB.AutoMigrate(&models.Post{})
	initializers.DB.AutoMigrate(&models.Like{})
	initializers.DB.AutoMigrate(&models.Comments{})
}
