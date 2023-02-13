package main

import "github.com/gin-gonic/gin"

func main() {
	// Create a Gin router
	router := gin.Default()

	// Define a route with a handler function
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})

	// Start the server
	router.Run(":8080")

}
