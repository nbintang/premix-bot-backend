package transaction

import "github.com/gin-gonic/gin"

func RegisterTransactionRoutes(rg *gin.RouterGroup, h *TransactionHandlerImpl) {
	transactionGroup := rg.Group("/transactions")
	transactionGroup.GET("", h.GetTransactions)
}
