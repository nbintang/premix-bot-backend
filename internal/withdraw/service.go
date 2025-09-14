package withdraw

import "premix-backend/internal/models"

type WithDrawServiceImpl struct {
	repo WithdrawRepository
}

func NewService(repo WithdrawRepository) WithdrawService {
	return &WithDrawServiceImpl{repo: repo}
}

func (s *WithDrawServiceImpl) GetWithdraws() ([]models.WithdrawRequest, error) {
	return s.repo.GetAllWithdraws()
}
