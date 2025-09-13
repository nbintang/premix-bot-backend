package withdraw

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Options(
	// fx.Provide(NewAuthService),  // nanti kalau ada service
	// fx.Provide(NewAuthHandler),  // nanti kalau ada handler
	fx.Invoke(register),
)

// register dipanggil otomatis sama fx saat startup
func register(db *gorm.DB, log *logrus.Logger) {
	// jalanin migration

}
