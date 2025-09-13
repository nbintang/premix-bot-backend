package models

import "time"

// User represents a user in the system
type User struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Email        string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
	Name         *string
	Saldo        float64 `gorm:"type:decimal(18,2);default:0.00"`
	TelegramID   *string `gorm:"uniqueIndex"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// Relations
	Bots             []Bot             `gorm:"constraint:OnDelete:CASCADE"`
	Transactions     []Transaction     `gorm:"constraint:OnDelete:CASCADE"`
	WithdrawRequests []WithdrawRequest `gorm:"constraint:OnDelete:CASCADE"`
	UserSessions     []UserSession     `gorm:"constraint:OnDelete:CASCADE"`
}

// UserSession represents refresh tokens or sessions
type UserSession struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	UserID    uint   `gorm:"not null"`
	Token     string `gorm:"uniqueIndex;not null"` //  refresh token harus base64 wir
	IP        *string
	UserAgent *string
	ExpiresAt time.Time `gorm:"not null"`
	Revoked   bool      `gorm:"default:false"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// Relations
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
