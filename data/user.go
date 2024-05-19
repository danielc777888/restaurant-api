package data

import (
	"github.com/google/uuid"
)

type User struct {
	ID              uuid.UUID `gorm:"type:uuid"`
	Name            string
	EmailAddress    string `gorm:"unique"`
	Password        string
	Locked          bool
	LoginAttempts   uint
	UserPermissions []UserPermission
}

type UserPermission struct {
	ID           uuid.UUID `gorm:"type:uuid"`
	UserID       uuid.UUID `gorm:"type:uuid"`
	PermissionID uuid.UUID `gorm:"type:uuid"`
	Permission   Permission
	RestaurantID *uuid.UUID `gorm:"type:uuid"`
}
