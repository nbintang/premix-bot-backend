package migration

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Invoke(AutoMigrate),
)
