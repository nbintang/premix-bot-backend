package migration

import (
	"premix-backend/internal/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// AutoMigrate runs Gorm auto migrations
func AutoMigrate(db *gorm.DB, log *logrus.Logger) error {
	models := []any{
		&models.User{},
		&models.UserSession{},
		&models.Bot{},
		&models.BotSubscription{},
		&models.Transaction{},
		&models.WithdrawRequest{},
	}

	for _, m := range models {
		if err := db.AutoMigrate(m); err != nil {
			log.Errorf("migration failed for %T: %v", m, err)
			return err
		}
	}

	log.Info("all migrations ran successfully")
	return nil
}
