package user

import "premix-backend/internal/models"

type UserServiceImpl struct {
	repo UserRepository
}

func NewService(repo UserRepository) UserService {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) GetUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}
