package user

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"size:100;not null;unique"`
	Email     string    `gorm:"size:100;not null;unique"`
	Password  string    `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}