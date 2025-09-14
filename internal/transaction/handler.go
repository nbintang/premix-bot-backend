package transaction

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandlerImpl struct {
	service TransactionService
}

func NewTransactionHandler(service TransactionService) *TransactionHandlerImpl {
	return &TransactionHandlerImpl{service: service}
}

func (h *TransactionHandlerImpl) GetTransactions(c *gin.Context) {
	transactions, err := h.service.GetTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}
