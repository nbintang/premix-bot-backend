package payment


type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetPayments() ([]Payment, error) {
	return s.repo.GetAllPayments()
}