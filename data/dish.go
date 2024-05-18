package data

import "github.com/google/uuid"

type Dish struct {
	ID           uuid.UUID `gorm:"type:uuid"`
	Name         string
	Description  string
	Price        uint
	Ratings      []Rating
	RestaurantID uuid.UUID `gorm:"index"`
}
