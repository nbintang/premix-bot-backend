package payment

import "premix-backend/internal/models"

type Repository interface {
	GetAllPayments() ([]models.Payment, error)
}

type Service interface {
	GetPayments() ([]models.Payment, error)
}