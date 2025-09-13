package models

import (
	"time"
)

// Bot represents a Telegram bot
type Bot struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	UserID      uint   `gorm:"not null"`
	BotToken    string `gorm:"not null"`
	BotUsername string `gorm:"not null"`
	IsActive    bool   `gorm:"default:true"`
	ExpiredAt   *time.Time

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// Relations
	User          User              `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Subscriptions []BotSubscription `gorm:"constraint:OnDelete:CASCADE"`
	Transactions  []Transaction     `gorm:"constraint:OnDelete:CASCADE"`
}

// BotSubscription represents a subscription for a bot
type BotSubscription struct {
	ID          uint    `gorm:"primaryKey;autoIncrement"`
	BotID       uint    `gorm:"not null"`
	Amount      float64 `gorm:"type:decimal(18,2);default:30000.00"`
	Status      string  `gorm:"not null"` // pending | paid | expired
	PeriodStart time.Time
	PeriodEnd   time.Time

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// Relations
	Bot Bot `gorm:"foreignKey:BotID;constraint:OnDelete:CASCADE"`
}
