package user



type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}