package config

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func StartServer(lc fx.Lifecycle, router *gin.Engine, log *logrus.Logger, cfg *EnvConfig) {
	srv := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: router,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Infof("server running on :%s 🚀", cfg.AppPort)
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.WithError(err).Fatal("failed to start server")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("stopping server...")
			return srv.Shutdown(ctx)
		},
	})
}
