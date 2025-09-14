package auth

import (
	"premix-backend/internal/auth/dto"
	"premix-backend/internal/shared/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes for register routes
func RegisterAuthRoutes(rg *gin.RouterGroup, h *HandlerImpl, v *middleware.Validation) {
	authGroup := rg.Group("/auth")

	// attach middleware ke route
	authGroup.POST("/login", v.Middleware(&dto.LoginRequest{}), h.Login)
	authGroup.POST("/register", v.Middleware(&dto.RegisterRequest{}), h.Register)
}
