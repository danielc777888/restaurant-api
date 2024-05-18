package data

import (
	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid"`
	Name          string
	EmailAddress  string `gorm:"unique"`
	Password      string
	Locked        bool
	LoginAttempts uint
	Permissions   []Permission
}
