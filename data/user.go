package data

import (
	"time"
)

type User struct {
	ID             uint
	Name           string
	EmailAddress   string `gorm:"unique"`
	Password       string
	Locked         bool
	LoginAttempts  uint
	Token          *string
	TokenCreatedAt *time.Time
}
