package main

import (
	"todo_api/config"
	"todo_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.InitDB(); err != nil {
		panic(err)
	}

	router := gin.Default()

	todoGroup := router.Group("/todos")
	routes.SetupTodoRoutes(todoGroup)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
