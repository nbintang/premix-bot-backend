package payment

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Options(
	fx.Provide(NewRepository),
	fx.Provide(NewService),
	fx.Provide(NewHandler),
	fx.Invoke(migrate),
)

func migrate(db *gorm.DB, log *logrus.Logger) { 
	RunMigration(db, log)
}
