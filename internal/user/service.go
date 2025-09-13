package user

import "premix-backend/internal/models"

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}
