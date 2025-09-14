// internal/app/module.go
package app

import (
	"premix-backend/internal/auth"
	"premix-backend/internal/bot"
	"premix-backend/internal/payment"
	"premix-backend/internal/transaction"
	"premix-backend/internal/user"
	"premix-backend/internal/withdraw"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var AppModule = fx.Options(
	auth.AuthModule,
	user.UserModule,
	payment.PaymentModule,
	bot.BotModule,
	transaction.TransactionModule,
	withdraw.WithdrawModule,
	fx.Invoke(bootstrap),
)

func bootstrap(
	router *gin.Engine,
	ah *auth.AuthHandlerImpl,
	uh *user.UserHandlerImpl,
	ph *payment.PaymentHandlerImpl,
	bh *bot.BotHandlerImpl,
	th *transaction.TransactionHandlerImpl,
	wh *withdraw.WithdrawHandlerImpl,
	log *logrus.Logger,
) {
	v1 := router.Group("/api/v1")
	auth.RegisterAuthRoutes(v1, ah)
	user.RegisterUserRoutes(v1, uh)
	payment.RegisterPaymentRoutes(v1, ph)
	bot.RegisterBotRoutes(v1, bh)
	transaction.RegisterTransactionRoutes(v1, th)
	withdraw.RegisterWithdrawRoutes(v1, wh)
	log.Info("routes registered successfully")
}
