package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"project/config"
	"project/models"
	"project/utils"
)

func Register(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	config.DB.Create(&user)

	token, _ := utils.GenerateToken(user.Username)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func Protected(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"token": "accessed protected route",
	})
}

func Login(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	var existingUser models.User
	config.DB.Where("username = ?", user.Username).First(&existingUser)

	if existingUser.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	token, _ := utils.GenerateToken(user.Username)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
