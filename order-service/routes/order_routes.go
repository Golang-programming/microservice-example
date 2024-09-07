package routes

import (
	"order-service/controllers"
	"order-service/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	orderGroup := router.Group("/order")
	{
		orderGroup.Use(middlewares.AuthMiddleware())
		orderGroup.POST("/", controllers.CreateOrder)
		orderGroup.GET("/:id", controllers.GetOrder)
		orderGroup.GET("/", controllers.GetAllOrders)
	}
}
