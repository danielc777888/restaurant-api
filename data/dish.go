package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Dish struct {
	ID           uuid.UUID `gorm:"type:uuid"`
	Name         string
	Description  string
	Price        uint
	Ratings      []Rating
	RestaurantID uuid.UUID `gorm:"index"`
}

type DishData struct {
	Db *gorm.DB
}

func NewDishData(Db *gorm.DB) *DishData {
	return &DishData{Db: Db}
}

// Create dish, inserts into database.
func (dishData *DishData) CreateDish(dish Dish) (*Dish, error) {
	result := dishData.Db.Create(&dish)
	if result.Error != nil {
		return nil, result.Error
	}
	return &dish, nil
}

// Update dish.
func (dishData *DishData) UpdateDish(restaurantID uuid.UUID, dish Dish) (*Dish, error) {
	result := dishData.Db.Where("id = ? AND restaurant_id = ?", dish.ID, restaurantID).Save(&dish)
	if result.Error != nil {
		return nil, result.Error
	}
	return &dish, nil
}

// Delete dish.
func (dishData *DishData) DeleteDish(restaurantID uuid.UUID, dishID uuid.UUID) error {
	result := dishData.Db.Where("id = ? AND restaurant_id = ?", dishID, restaurantID).Delete(&Dish{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Get dish.
func (dishData *DishData) GetDish(restaurantID uuid.UUID, dishID uuid.UUID, dish Dish) (*Dish, error) {
	result := dishData.Db.Where("id = ? AND restaurant_id = ?", dishID, restaurantID).First(&dish)
	if result.Error != nil {
		return nil, result.Error
	}
	return &dish, nil
}

// List dishes.
func (dishData *DishData) ListDishes(restaurantID uuid.UUID) ([]Dish, error) {
	var dishes []Dish
	result := dishData.Db.Where("restaurant_id = ?", restaurantID).Find(&dishes)
	if result.Error != nil {
		return nil, result.Error
	}
	return dishes, nil
}
