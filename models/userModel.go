package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id              int            `json:"id" gorm:"type:INT(10); NOT NULL AUTOINCREMENT;primaryKey"`
	FirstName       string         `json:"first_name"`
	LastName        string         `json:"last_name"`
	Email           string         `json:"email" gorm:"unique"`
	Password        string         `json:"password"`
	IsEmailVerified bool           `json:"is_email_verified" gorm:"default:false"`
	PhoneNumber     string         `json:"phone_number" gorm:"unique"`
	CreatedAt       time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Verify_Email struct {
	// gorm.Model
	Id         int       `json:"id" gorm:"type:INT(10); NOT NULL AUTOINCREMENT;primaryKey"`
	Email      string    `json:"email" gorm:"NOT NULL;"`
	SecretCode string    `json:"secret_code" gorm:"NOT NULL;"`
	IsUsed     bool      `json:"is_used" gorm:" NOT NULL, default:false"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	ExpiredAt  time.Time `json:"expired_at" gorm:"NOT NULL"`
	UserId     int       `json:"userId" gorm:"NOT NULL;"`
	User       User      `gorm:"constraint:OnDelete:CASCADE"`
}
