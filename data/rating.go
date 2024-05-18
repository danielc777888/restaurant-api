package data

import (
	"database/sql"

	"github.com/google/uuid"
)

type Rating struct {
	ID           uuid.UUID `gorm:"type:uuid"`
	Description  string
	Sentiment    sql.NullBool
	DishID       uuid.UUID `gorm:"index"`
	RestaurantID uuid.UUID `gorm:"index"`
}
