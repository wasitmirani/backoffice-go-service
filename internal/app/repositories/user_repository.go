package repositories

import (
	"errors"
	"backendapp/internal/app/models"
	"backendapp/db"
)



func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}


func FindUserByUUID(uuid string) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}
