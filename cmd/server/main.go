package main

import (
	"premix-backend/internal/app"
	"premix-backend/internal/shared/config"
	"premix-backend/internal/shared/middleware"
	"premix-backend/internal/shared/migration"

	"go.uber.org/fx"
)

func main() {
	appServer := fx.New(
		config.Module, // config tetap dipisah
		middleware.Module,
		migration.Module,
		app.Module, // semua module bisnis
	)
	appServer.Run()
}
