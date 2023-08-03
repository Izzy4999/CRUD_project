package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id          int `gorm:"type:INT(10) UNSIGNED NOT NULL AUTOINCREMENT;primaryKey"`
	FirstName   string
	LastName    string
	Email       string `gorm:"unique"`
	Password    string
	PhoneNumber string         `gorm:"unique"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
