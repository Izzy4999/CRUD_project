package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id          int            `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTOINCREMENT;primaryKey"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	Email       string         `json:"email" gorm:"unique"`
	Password    string         `json:"password"`
	PhoneNumber string         `json:"phone_number" gorm:"unique"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
