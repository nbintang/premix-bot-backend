package user

import "premix-backend/internal/models"

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
}

type UserService interface {
	GetUsers() ([]models.User, error)
}
