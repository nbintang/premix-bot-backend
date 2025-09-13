package user

import "premix-backend/internal/models"

type Repository interface {
	GetAllUsers() ([]models.User, error)
}

type Service interface {
	GetUsers() ([]models.User, error)
}
