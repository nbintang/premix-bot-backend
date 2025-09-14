package bot

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BotHandlerImpl struct {
	service BotService
}

func NewBotHandler(service BotService) *BotHandlerImpl {
	return &BotHandlerImpl{service: service}
}

func (h *BotHandlerImpl) GetBots(c *gin.Context) {
	bots, err := h.service.GetBots()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bots)
}
