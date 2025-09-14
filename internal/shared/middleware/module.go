package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

// Module export module
var Module = fx.Options(
	fx.Provide(NewValidation),
	fx.Invoke(RegisterMiddlewares),
)

// RegisterMiddlewares inject middleware ke Gin engine
func RegisterMiddlewares(router *gin.Engine, logger *logrus.Logger) {
	router.Use(TraceIDMiddleware(logger))
}
