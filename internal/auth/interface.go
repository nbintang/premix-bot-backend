package auth

// Service interface
type Service interface {
	Login(username string, password string) (string, error)
	Register(username string, name string, password string, confirmPassword string) (string, error)
}
