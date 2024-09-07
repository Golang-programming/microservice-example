package services

import (
	"product-service/models"
	"product-service/repositories"
)

type ProductService struct {
	ProductRepository repositories.ProductRepository
}

type CreateProductDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return ProductService{ProductRepository: repo}
}

func (s *ProductService) CreateProduct(dto CreateProductDTO) (models.Product, error) {
	product := models.Product{Name: dto.Name, Description: dto.Description, Price: dto.Price}
	return s.ProductRepository.Save(product)
}

func (s *ProductService) GetProduct(id string) (models.Product, error) {
	return s.ProductRepository.FindByID(id)
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.ProductRepository.FindAll()
}
