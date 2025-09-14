package withdraw

import (
	"go.uber.org/fx"
)

// WithdrawModule export module
var WithdrawModule = fx.Options(
	fx.Provide(NewWithdrawRepository), // nanti kalau ada service
	fx.Provide(NewService),            // nanti kalau ada handler
	fx.Provide(NewHandler),            // nanti kalau ada handler

)
