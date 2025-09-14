package payment

import (
	"go.uber.org/fx"
)

var PaymentModule = fx.Options(
	fx.Provide(NewPaymentRepository),
	fx.Provide(NewPaymentService),
	fx.Provide(NewPaymentHandler),
)