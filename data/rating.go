package data

import (
	"github.com/google/uuid"
)

type Rating struct {
	ID           uuid.UUID `gorm:"type:uuid"`
	Description  string
	Sentiment    *string
	DishID       uuid.UUID `gorm:"index"`
	RestaurantID uuid.UUID `gorm:"index"`
}
