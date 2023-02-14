```go

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// code before processing the request
		fmt.Println("nice")
		c.Next()
		// code after processing the request
	}
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// Apply the middleware only for this route


	// private := r.Group("/private")
	// private.Use(LoggerMiddleware())


	private := r.Group("/private", LoggerMiddleware())

	private.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from a private route!",
		})
	})

	// Apply the middleware only for this route
	r.GET("/single", LoggerMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from a single route!",
		})
	})

	r.Run()
}

```
