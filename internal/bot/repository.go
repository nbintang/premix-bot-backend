package bot

import (
	"premix-backend/internal/models"

	"gorm.io/gorm"
)

type BotRepositoryImpl struct {
	db *gorm.DB
}

func NewBotRepository(db *gorm.DB) BotRepository {
	return &BotRepositoryImpl{db: db}
}

func (r *BotRepositoryImpl) GetAllBots() ([]models.Bot, error) {
	var bots []models.Bot
	if err := r.db.Find(&bots).Error; err != nil {
		return nil, err
	}
	return bots, nil
}
