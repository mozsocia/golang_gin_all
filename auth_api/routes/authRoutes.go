package routes

import (
	"github.com/gin-gonic/gin"

	"project/controllers"
)

func AuthRoutes(route *gin.Engine) {
	auth := route.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}
}
