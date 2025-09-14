package payment

import (
	"premix-backend/internal/models"

	"gorm.io/gorm"
)

type PaymentRepositoryImpl struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &PaymentRepositoryImpl{db: db}
}

func (r *PaymentRepositoryImpl) GetAllPaymentMethods() ([]models.Payment, error) {
	var paymentMethods []models.Payment
	if err := r.db.Find(&paymentMethods).Error; err != nil {
		return nil, err
	}
	return paymentMethods, nil
}
