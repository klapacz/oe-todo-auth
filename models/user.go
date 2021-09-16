package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string `gorm:"type:uuid;default:uuid_generate_v4()"`
	Password string
	Email    string
}
