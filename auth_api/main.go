package main

import (
	"github.com/gin-gonic/gin"

	"project/config"
	"project/controllers"
	"project/middlewares"
	"project/routes"
)

func main() {
	r := gin.Default()

	// Load environment variables from .env file
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	config.ConnectDB()

	// Initialize routes
	routes.AuthRoutes(r)
	// Protected route (authentication required)
	r.GET("/protected", middlewares.AuthMiddleware(), controllers.Protected)

	r.Run(":8080")
}
