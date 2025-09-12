package user

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Module membungkus semua komponen User ke dalam fx.Options
var Module = fx.Options(
	fx.Provide(NewRepository),
	fx.Provide(NewService),
	fx.Provide(NewHandler),
	fx.Invoke(migrate),
)

func migrate(db *gorm.DB, log *logrus.Logger) {
    RunMigration(db, log)
}
