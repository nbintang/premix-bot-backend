package payment

import "premix-backend/internal/models"

type PaymentRepository interface {
	GetAllPaymentMethods() ([]models.Payment, error)
}

type PaymentService interface {
	GetPaymentMethods() ([]models.Payment, error)
}
