package user

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(rg *gin.RouterGroup, h *UserHandlerImpl) {
	userGroup := rg.Group("/users")
	userGroup.GET("", h.GetUsers)
}
