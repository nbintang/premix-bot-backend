package payment

import "premix-backend/internal/models"

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetPayments() ([]models.Payment, error) {
	return s.repo.GetAllPayments()
}