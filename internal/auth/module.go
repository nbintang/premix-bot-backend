package auth

import (
	"go.uber.org/fx"
)

// AuthModule auth module
var AuthModule = fx.Options(
	fx.Provide(NewService), // nanti kalau ada service
	fx.Provide(NewHandler), // nanti kalau ada handler
)
