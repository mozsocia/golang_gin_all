package config

import (
	"project/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	// Automigrate the models
	db.AutoMigrate(&models.User{})

	DB = db

	return nil
}
