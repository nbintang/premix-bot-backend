package config

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func NewRouter() *gin.Engine {
    return gin.Default()
}

func StartServer(lc fx.Lifecycle, router *gin.Engine, log *logrus.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Info("server running on :8080 🚀")
				if err := router.Run(":8080"); err != nil {
					log.WithError(err).Fatal("failed to start server")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("server stopped")
			return nil
		},
	})
}