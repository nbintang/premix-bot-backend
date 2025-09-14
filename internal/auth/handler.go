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

func (h *AuthHandlerImpl) Register(c *gin.Context) {
	username := c.PostForm("username")
	name := c.PostForm("name")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirmPassword")

	token, err := h.service.Register(username, name, password, confirmPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
