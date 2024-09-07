package repositories

import (
	"order-service/database"
	"order-service/models"
)

func Save(order models.Order) (models.Order, error) {
	if err := database.DB.Create(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}

func FindByID(id string) (models.Order, error) {
	var order models.Order
	if err := database.DB.First(&order, id).Error; err != nil {
		return order, err
	}
	return order, nil
}

func FindAll() ([]models.Order, error) {
	var orders []models.Order
	if err := database.DB.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
