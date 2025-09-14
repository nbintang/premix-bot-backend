package auth

type AuthServiceImpl struct{}

func NewAuthService() AuthService {
	return &AuthServiceImpl{}
}

func (s *AuthServiceImpl) Login(username string, password string) (string, error) {
	return "dummy-token", nil
}

func (s *AuthServiceImpl) Register(username string, name string, password string, confirmPassword string) (string, error) {
	return "dummy-tokem", nil
}
