package main

import (
	"todoapi/models"
	"todoapi/routes"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	// "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.Todo{})

	r := gin.Default()
	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})
	routes.TodoRoutes(r, db)
	r.Run(":8080")
}
