package services

import (
	"errors"
	"backendapp/internal/app/models"
	"backendapp/internal/app/repositories"
	"backendapp/internal/pkg/utils"
)

func Login(email, password string) (*models.User, error) {
	user, err := repositories.FindUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
