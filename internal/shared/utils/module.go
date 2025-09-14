package utils

import (
	"go.uber.org/fx"
)

// Module export module
var Module = fx.Options(
	fx.Provide(NewResponseHelper),
)
