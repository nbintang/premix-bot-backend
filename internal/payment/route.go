package payment

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
    paymentGroup := rg.Group("/payments")
    paymentGroup.GET("/", h.GetUsers)
}