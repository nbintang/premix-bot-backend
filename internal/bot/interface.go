package bot

import "premix-backend/internal/models"

type BotRepository interface {
	GetAllBots() ([]models.Bot, error)
}

type BotService interface {
	GetBots() ([]models.Bot, error)
}
