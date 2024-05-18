package data

import (
	"github.com/google/uuid"
)

type Permission struct {
	ID     uuid.UUID `gorm:"type:uuid"`
	Name   string
	UserID uuid.UUID `gorm:"type:uuid"`
}
