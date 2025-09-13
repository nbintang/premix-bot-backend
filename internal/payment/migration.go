package payment

import (
	"premix-backend/internal/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB, log *logrus.Logger) {
	if err := db.AutoMigrate(&models.Payment{}); err != nil {
		log.WithError(err).Fatal("failed to migrate payment table")
	} else {
		log.Info("payment table migrated successfully")
	}
}
