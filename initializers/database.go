package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dbURI := os.Getenv("DB_URI")

	DB, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		panic("Database connection error")
	}
}
