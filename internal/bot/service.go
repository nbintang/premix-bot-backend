package bot

import "premix-backend/internal/models"

type BotServiceImpl struct {
	repo BotRepository
}

func NewBotService(repo BotRepository) BotService {
	return &BotServiceImpl{repo: repo}
}

func (s *BotServiceImpl) GetBots() ([]models.Bot, error) {
	return s.repo.GetAllBots()
}
