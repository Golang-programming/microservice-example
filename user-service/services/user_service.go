package services

import (
	"user-service/models"
	"user-service/repositories"
)

type CreateUserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(dto CreateUserDTO) (models.User, error) {
	user := models.User{Name: dto.Name, Email: dto.Email}
	return repositories.Save(user)
}

func GetUser(id string) (models.User, error) {
	return repositories.FindByID(id)
}

func GetAllUsers() ([]models.User, error) {
	return repositories.FindAll()
}
