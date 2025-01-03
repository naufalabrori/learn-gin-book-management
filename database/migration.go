package database

import (
	"learn-go-gin/config"
	"learn-go-gin/models"
)

func Migrate() {
	config.DB.AutoMigrate(&models.User{}, &models.Book{}, &models.Transaction{}, &models.Category{}, &models.Fines{})
}
