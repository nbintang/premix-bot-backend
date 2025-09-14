package auth

import "github.com/gin-gonic/gin"

func RegisterAuthRoutes(rg *gin.RouterGroup, h *AuthHandlerImpl) {
	authGroup := rg.Group("/auth")
	authGroup.POST("/login", h.Login)
	authGroup.POST("/register", h.Register)
}
