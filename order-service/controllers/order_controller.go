package controllers

import (
	"net/http"
	"order-service/services"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var orderDto services.CreateOrderDTO
	if err := c.ShouldBindJSON(&orderDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order, err := services.CreateOrder(orderDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

func GetOrder(c *gin.Context) {
	id := c.Param("id")
	order, err := services.GetOrder(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

func GetAllOrders(c *gin.Context) {
	orders, err := services.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}
