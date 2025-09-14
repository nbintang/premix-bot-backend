package withdraw

import "premix-backend/internal/models"

type WithdrawRepository interface {
	GetAllWithdraws() ([]models.WithdrawRequest, error)
}

type WithdrawService interface {
	GetWithdraws() ([]models.WithdrawRequest, error)
}
