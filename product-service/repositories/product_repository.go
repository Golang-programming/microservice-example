package repositories

import (
	"product-service/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{DB: db}
}

func (r *ProductRepository) Save(product models.Product) (models.Product, error) {
	if err := r.DB.Create(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (r *ProductRepository) FindByID(id string) (models.Product, error) {
	var product models.Product
	if err := r.DB.First(&product, id).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (r *ProductRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	if err := r.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
