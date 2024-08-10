package main

import (
	"transaction_system/config"
	"transaction_system/internal/server"
	"transaction_system/internal/server/handler"
	"transaction_system/pkg/transaction"
	"transaction_system/utils/initialize"

	"go.uber.org/fx"
)

func serverRun() {
	app := fx.New(
		fx.Provide(
			// postgresql
			initialize.NewDB,
		),
		config.Module,
		initialize.Module,
		server.Module,
		handler.Module,
		transaction.Module,
	)

	// Run app forever
	app.Run()
}
