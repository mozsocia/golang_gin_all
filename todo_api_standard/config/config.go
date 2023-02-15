package config

import (
	"todo_api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	// Migrate the schema
	DB.AutoMigrate(&models.Todo{})
	return nil
}
