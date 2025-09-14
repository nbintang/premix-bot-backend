package withdraw

import "github.com/gin-gonic/gin"

func RegisterWithdrawRoutes(rg *gin.RouterGroup, h *WithdrawHandlerImpl) {
	withdrawGroup := rg.Group("/withdraws")
	withdrawGroup.GET("", h.GetUsers)
}
