package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandlerImpl struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandlerImpl {
	return &UserHandlerImpl{service: service}
}

func (h *UserHandlerImpl) GetUsers(c *gin.Context) {
	users, err := h.service.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
