package withdraw

import (
	"premix-backend/internal/models"

	"gorm.io/gorm"
)

type WithdrawRepositoryImpl struct {
	db *gorm.DB
}

func NewWithdrawRepository(db *gorm.DB) WithdrawRepository {
	return &WithdrawRepositoryImpl{db: db}
}

func (r *WithdrawRepositoryImpl) GetAllWithdraws() ([]models.WithdrawRequest, error) {
	var withdraws []models.WithdrawRequest
	if err := r.db.Find(&withdraws).Error; err != nil {
		return nil, err
	}
	return withdraws, nil
}
