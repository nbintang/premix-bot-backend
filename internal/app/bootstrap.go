// internal/app/module.go
package app

import (
	"premix-backend/internal/payment"
	"premix-backend/internal/user"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Options(
	user.Module,
	payment.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(router *gin.Engine, uh *user.Handler, ph *payment.Handler, log *logrus.Logger) {

	v1 := router.Group("/api/v1")

	user.RegisterRoutes(v1, uh)
	payment.RegisterRoutes(v1, ph)

	log.Info("routes registered successfully")
}
