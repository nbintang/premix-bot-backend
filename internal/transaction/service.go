package transaction

import "premix-backend/internal/models"

type TransactionServiceImpl struct {
	repo TransactionRepository
}

func NewTransactionService(repo TransactionRepository) TransactionService {
	return &TransactionServiceImpl{repo: repo}
}

func (s *TransactionServiceImpl) GetTransactions() ([]models.Transaction, error) {
	return s.repo.GetAllTransactions()
}
