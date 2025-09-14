package auth

// ServiceImpl struct
type ServiceImpl struct{}

// NewService init
func NewService() Service {
	return &ServiceImpl{}
}

// Login for login
func (s *ServiceImpl) Login(username string, password string) (string, error) {
	return "dummy-token", nil
}

// Register for register
func (s *ServiceImpl) Register(username string, name string, password string, confirmPassword string) (string, error) {
	return "dummy-tokem", nil
}
