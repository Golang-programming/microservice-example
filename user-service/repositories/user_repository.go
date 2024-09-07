package repositories

import (
	"user-service/database"
	"user-service/models"
)

func Save(user models.User) (models.User, error) {
	if err := database.DB.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func FindByID(id string) (models.User, error) {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

// FindAll retrieves all users
func FindAll() ([]models.User, error) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
