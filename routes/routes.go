package routes

import (
	"login-api/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/login", controllers.Login)
		userRoutes.POST("/logout", controllers.Logout)
	}
}
