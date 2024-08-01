package database

import (
	"fmt"
	"golang_blog/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Connect() {
	cfg := config.Get()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connot connect to DB")
		return
	}
	DB = db
}
