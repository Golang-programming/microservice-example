package controllers

import (
	"net/http"
	"product-service/services"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService services.ProductService
}

func NewProductController(service services.ProductService) ProductController {
	return ProductController{ProductService: service}
}

func (ctrl *ProductController) CreateProduct(c *gin.Context) {
	var productDto services.CreateProductDTO
	if err := c.ShouldBindJSON(&productDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := ctrl.ProductService.CreateProduct(productDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (ctrl *ProductController) GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := ctrl.ProductService.GetProduct(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (ctrl *ProductController) GetAllProducts(c *gin.Context) {
	products, err := ctrl.ProductService.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}
