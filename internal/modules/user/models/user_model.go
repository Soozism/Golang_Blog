package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"verchar:191"`
	Email    string `gorm:"verchar:191"`
	Password string `gorm:"verchar:191"`
}
