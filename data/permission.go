package data

import (
	"github.com/google/uuid"
)

type Permission struct {
	ID  uuid.UUID `gorm:"type:uuid"`
	Key string
}
