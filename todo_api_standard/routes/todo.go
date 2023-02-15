package routes

import (
	"github.com/gin-gonic/gin"

	"todo_api/controllers"
)

func SetupTodoRoutes(todoGroup *gin.RouterGroup) {
	// todoGroup := router.Group("/todos")
	todoGroup.GET("/", controllers.GetAllTodos)
	todoGroup.GET("/:id", controllers.GetTodo)
	todoGroup.POST("/", controllers.CreateTodo)
	todoGroup.PUT("/:id", controllers.UpdateTodo)
	todoGroup.DELETE("/:id", controllers.DeleteTodo)
}
