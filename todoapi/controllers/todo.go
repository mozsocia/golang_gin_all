package controllers

import (
	"todoapi/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type TodoController struct {
	DB *gorm.DB
}

func (tc TodoController) GetTodos(c *gin.Context) {
	var todos []models.Todo
	tc.DB.Find(&todos)

	c.JSON(200, todos)
}

func (tc TodoController) GetTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.Todo
	tc.DB.First(&todo, id)

	if todo.ID == 0 {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(200, todo)
}

func (tc TodoController) CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": "Please provide all the required data"})
		return
	}

	tc.DB.Create(&todo)
	c.JSON(201, todo)
}

func (tc TodoController) UpdateTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.Todo
	tc.DB.First(&todo, id)

	if todo.ID == 0 {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}

	if err := c.BindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": "Please provide all the required data"})
		return
	}

	tc.DB.Save(&todo)
	c.JSON(200, todo)
}

func (tc TodoController) DeleteTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.Todo
	tc.DB.First(&todo, id)

	if todo.ID == 0 {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}

	tc.DB.Delete(&todo)
	c.JSON(204, gin.H{"success": "Todo deleted"})
}
