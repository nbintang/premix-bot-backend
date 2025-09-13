package user

import (
	"go.uber.org/fx"
)

// Module membungkus semua komponen User ke dalam fx.Options
var Module = fx.Options(
	fx.Provide(NewRepository),
	fx.Provide(NewService),
	fx.Provide(NewHandler),
)
