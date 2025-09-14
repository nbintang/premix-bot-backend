package withdraw

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WithdrawHandlerImpl struct {
	service WithdrawService
}

func NewHandler(service WithdrawService) *WithdrawHandlerImpl {
	return &WithdrawHandlerImpl{service: service}
}

func (h *WithdrawHandlerImpl) GetUsers(c *gin.Context) {
	withdraws, err := h.service.GetWithdraws()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, withdraws)
}
