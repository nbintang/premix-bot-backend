package config

import (
	"fmt"


	"github.com/sirupsen/logrus" 
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(cfg *EnvConfig, logger *logrus.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.WithError(err).Error("failed to connect to database")
		return nil, err
	}
	logger.Info("database connected successfully")
	return db, nil
}