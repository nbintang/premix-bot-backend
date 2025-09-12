package payment

type Repository interface {
	GetAllPayments() ([]Payment, error)
}

type Service interface {
	GetPayments() ([]Payment, error)
}