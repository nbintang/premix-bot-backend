package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
    userGroup := rg.Group("/users")
    userGroup.GET("", h.GetUsers)
}