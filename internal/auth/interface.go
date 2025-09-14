package auth

type AuthService interface {
	Login(username string, password string) (string, error)
	Register(username string, name string, password string, confirmPassword string) (string, error)
}
