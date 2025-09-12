package user

type Repository interface {
	GetAllUsers() ([]User, error)
}

type Service interface {
	GetUsers() ([]User, error)
}