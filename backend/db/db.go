package db

import (
	"log"
	"os"

	"url-shortener/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DATABASE_URL")

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB接続失敗", err)
	}

	DB.AutoMigrate(&models.User{}, &models.URL{})
	log.Println("DB接続成功")
}
