package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"todo_api/config"
	"todo_api/models"
)

// type CreateTodoInput struct {
// 	Title       string `json:"title" binding:"required"`
// 	Description string `json:"description"`
// }

// type UpdateTodoInput struct {
// 	Title       string `json:"title"`
// 	Description string `json:"description"`
// }

func GetAllTodos(c *gin.Context) {
	var todos []models.Todo
	if err := config.DB.Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving Todos"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func CreateTodo(c *gin.Context) {
	var input models.CreateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{Title: input.Title, Description: input.Description}
	if err := config.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating Todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func GetTodo(c *gin.Context) {
	var todo models.Todo

	if err := config.DB.First(&todo, c.Param("id")).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving Todo"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo

	if err := config.DB.First(&todo, c.Param("id")).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating Todo"})
		}
		return
	}

	var input models.UpdateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&todo).Updates(models.Todo{Title: input.Title, Description: input.Description}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating Todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo

	if err := config.DB.First(&todo, c.Param("id")).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting Todo"})
		}
		return
	}

	if err := config.DB.Delete(&todo, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting Todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Todo deleted"})
}
