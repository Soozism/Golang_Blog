package models

import (
	"golang_blog/internal/modules/user/models"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string `gorm:"verchar:191"`
	Content string `gorm:"text"`
	UserID  uint
	User    models.User
}
