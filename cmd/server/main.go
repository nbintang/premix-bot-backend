package main

import (
	"premix-backend/internal/app"
	"premix-backend/internal/shared/config"

	"go.uber.org/fx"
)

func main() {
	appServer := fx.New(
		config.Module, // config tetap dipisah
		app.Module,    // semua module bisnis
	)
	appServer.Run()
}
