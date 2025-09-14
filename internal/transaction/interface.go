package transaction

import "premix-backend/internal/models"

type TransactionRepository interface {
	GetAllTransactions() ([]models.Transaction, error)
}

type TransactionService interface {
	GetTransactions() ([]models.Transaction, error)
}
