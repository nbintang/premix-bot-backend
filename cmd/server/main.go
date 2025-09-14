package main

import (
	"premix-backend/internal/app"
	"premix-backend/internal/shared/config"
	"premix-backend/internal/shared/middleware"
	"premix-backend/internal/shared/migration"
	"premix-backend/internal/shared/utils"

	"go.uber.org/fx"
)

func main() {
	appServer := fx.New(
		config.Module, // config tetap dipisah
		middleware.Module,
		migration.Module,
		utils.Module,
		app.AppModule, // semua module bisnis
	)
	appServer.Run()
}
