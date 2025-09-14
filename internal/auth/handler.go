package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandlerImpl struct {
	service AuthService
}

func NewAuthHandler(service AuthService) *AuthHandlerImpl {
	return &AuthHandlerImpl{service: service}
}

func (h *AuthHandlerImpl) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The input must be provided!"})
		return
	}
	token, err := h.service.Login(username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, token)
}
