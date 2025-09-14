package auth

type AuthService interface {
	Login(username string, password string) (string, error)
}
