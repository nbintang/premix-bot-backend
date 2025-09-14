package payment

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentHandlerImpl struct {
	service PaymentService
}

func NewPaymentHandler(service PaymentService) *PaymentHandlerImpl {
	return &PaymentHandlerImpl{service: service}
}

func (h *PaymentHandlerImpl) GetUsers(c *gin.Context) {
	paymentMethods, err := h.service.GetPaymentMethods()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, paymentMethods)
}
