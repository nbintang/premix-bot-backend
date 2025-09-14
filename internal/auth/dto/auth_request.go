package dto

// LoginRequest body for login
type LoginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// RegisterRequest body for register
type RegisterRequest struct {
	Username        string `json:"username" form:"username" binding:"required"`
	Name            string `json:"name" form:"name" binding:"required"`
	Password        string `json:"password" form:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required,eqfield=Password"`
}
