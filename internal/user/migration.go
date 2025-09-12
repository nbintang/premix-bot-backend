package user

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB, log *logrus.Logger) {
	if err := db.AutoMigrate(&User{}); err != nil {
		log.WithError(err).Fatal("failed to migrate user table")
	} else {
		log.Info("user table migrated successfully")
	}
}
