package payment

import (
	"premix-backend/internal/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}


func (r *repository) GetAllPayments() ([]models.Payment, error) {
	var users []models.Payment
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
