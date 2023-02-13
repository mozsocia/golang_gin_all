package routes

import (
	"todoapi/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// "github.com/jinzhu/gorm"
)

func TodoRoutes(r *gin.Engine, db *gorm.DB) {
	todoController := controllers.TodoController{DB: db}

	r.GET("/todos", todoController.GetTodos)
	r.GET("/todos/:id", todoController.GetTodo)
	r.POST("/todos", todoController.CreateTodo)
	r.PUT("/todos/:id", todoController.UpdateTodo)
	r.DELETE("/todos/:id", todoController.DeleteTodo)
}
