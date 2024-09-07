package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID        uint `json:"id" gorm:"primaryKey"`
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
	Total     uint `json:"total"`
}
