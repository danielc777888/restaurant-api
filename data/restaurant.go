package data

import "github.com/google/uuid"

type Restaurant struct {
	ID     uuid.UUID `gorm:"type:uuid"`
	Name   string    `gorm:"unique"`
	Dishes []Dish
}
