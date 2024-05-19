package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Restaurant struct {
	ID     uuid.UUID `gorm:"type:uuid"`
	Name   string    `gorm:"unique"`
	Dishes []Dish
}

type RestaurantData struct {
	Db *gorm.DB
}

func NewRestaurantData(Db *gorm.DB) *RestaurantData {
	return &RestaurantData{Db: Db}
}

// List restaurants
func (restaurantData *RestaurantData) ListRestaurants() ([]Restaurant, error) {
	var restaurants []Restaurant
	result := restaurantData.Db.Find(&restaurants)
	if result.Error != nil {
		return nil, result.Error
	}
	return restaurants, nil
}
