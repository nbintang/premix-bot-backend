package user

import (
	"go.uber.org/fx"
)

// UserModule membungkus semua komponen User ke dalam fx.Options
var UserModule = fx.Options(
	fx.Provide(NewUserRepository),
	fx.Provide(NewService),
	fx.Provide(NewUserHandler),
)
