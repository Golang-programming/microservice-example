package routes

import (
	"user-service/controllers"
	"user-service/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	userGroup := router.Group("/")
	{
		userGroup.POST("/", controllers.CreateUser)
		userGroup.POST("/generate-api-key", controllers.GenerateAPIKey)
		userGroup.Use(middlewares.AuthMiddleware())
		userGroup.GET("/:id", controllers.GetUser)
		userGroup.GET("/", controllers.GetAllUsers)
	}
}
