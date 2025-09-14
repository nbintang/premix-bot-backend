package payment

import "github.com/gin-gonic/gin"

func RegisterPaymentRoutes(rg *gin.RouterGroup, h *PaymentHandlerImpl) {
	paymentGroup := rg.Group("/payments")
	paymentGroup.GET("/", h.GetUsers)
}
