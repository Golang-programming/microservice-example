package routes

import (
	"product-service/controllers"
	"product-service/middlewares"
	"product-service/repositories"
	"product-service/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	productRepository := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepository)
	productController := controllers.NewProductController(productService)

	productGroup := router.Group("/product")
	{
		productGroup.Use(middlewares.AuthMiddleware()) // Apply authentication middleware
		productGroup.POST("/", productController.CreateProduct)
		productGroup.GET("/:id", productController.GetProduct)
		productGroup.GET("/", productController.GetAllProducts)
	}
}
