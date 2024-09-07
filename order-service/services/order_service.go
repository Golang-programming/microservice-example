package services

import (
	"order-service/models"
	"order-service/repositories"
)

type CreateOrderDTO struct {
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

func CreateOrder(dto CreateOrderDTO) (models.Order, error) {
	order := models.Order{ProductID: dto.ProductID, Quantity: dto.Quantity, Total: dto.Quantity * 100} // Assuming price is 100 per unit
	return repositories.Save(order)
}

func GetOrder(id string) (models.Order, error) {
	return repositories.FindByID(id)
}

func GetAllOrders() ([]models.Order, error) {
	return repositories.FindAll()
}
