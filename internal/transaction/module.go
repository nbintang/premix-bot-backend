package transaction

import (
	"go.uber.org/fx"
)

var TransactionModule = fx.Options(
	fx.Provide(NewTransactionRepository),
	fx.Provide(NewTransactionService),  // nanti kalau ada service
	fx.Provide(NewTransactionHandler),  // nanti kalau ada handler
)
