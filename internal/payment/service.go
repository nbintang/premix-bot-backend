package payment

import "premix-backend/internal/models"

type PaymentServiceImpl struct {
	repo PaymentRepository
}

func NewPaymentService(repo PaymentRepository) PaymentService {
	return &PaymentServiceImpl{repo: repo}
}

func (s *PaymentServiceImpl) GetPaymentMethods() ([]models.Payment, error) {
	return s.repo.GetAllPaymentMethods()
}
