package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Invoke(RegisterMiddlewares),
)

// RegisterMiddlewares inject middleware ke Gin engine
func RegisterMiddlewares(router *gin.Engine, logger *logrus.Logger) {
	router.Use(TraceIDMiddleware(logger))
}
