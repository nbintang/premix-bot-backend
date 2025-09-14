package bot

import (
	"go.uber.org/fx"
)

var BotModule = fx.Options(
	fx.Provide(NewBotRepository),
	fx.Provide(NewBotService), // nanti kalau ada service
	fx.Provide(NewBotHandler), // nanti kalau ada handler
)
