package auth

import (
	"go.uber.org/fx"
)

var AuthModule = fx.Options(
	fx.Provide(NewAuthService),  // nanti kalau ada service
	fx.Provide(NewAuthHandler),  // nanti kalau ada handler
)
