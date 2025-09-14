package bot

import "github.com/gin-gonic/gin"

func RegisterBotRoutes(rg *gin.RouterGroup, h *BotHandlerImpl) {
	botGroup := rg.Group("/bots")
	botGroup.GET("/", h.GetBots)
}
