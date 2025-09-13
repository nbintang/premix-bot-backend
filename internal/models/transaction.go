package models

import (
	"time"
)

// Transaction represents a payment or subscription transaction
type Transaction struct {
	ID     uint    `gorm:"primaryKey;autoIncrement"`
	BotID  uint    `gorm:"not null"`
	UserID uint    `gorm:"not null"`
	Type   string  `gorm:"not null"` // income | withdraw | subscription
	Amount float64 `gorm:"type:decimal(18,2);not null"`
	Status string  `gorm:"not null"` // pending | success | failed
	RefID  *string
	Note   *string

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// Relations
	Bot  Bot  `gorm:"foreignKey:BotID;constraint:OnDelete:CASCADE"`
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

// WithdrawRequest represents a withdrawal request
type WithdrawRequest struct {
	ID          uint    `gorm:"primaryKey;autoIncrement"`
	UserID      uint    `gorm:"not null"`
	Amount      float64 `gorm:"type:decimal(18,2);not null"`
	BankName    *string
	BankAccount *string
	Status      string `gorm:"not null"` // pending | approved | rejected

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
