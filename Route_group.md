```go
package main

import (
	"github.com/gin-gonic/gin"
)

func SetupAdminRoutes(admin *gin.RouterGroup) {
	admin.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the admin dashboard",
		})
	})

	admin.GET("/profile", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "You are viewing your profile",
		})
	})
}

func main() {
	r := gin.Default()

	// Define a route group for "admin" routes

	admin := r.Group("/admin")
	SetupAdminRoutes(admin)

	SetupAdminRoutes(r.Group("/admin2"))

	// Start the server
	r.Run()
}
```
