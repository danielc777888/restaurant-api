package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rating struct {
	ID           uuid.UUID `gorm:"type:uuid"`
	Description  string
	Sentiment    *string
	DishID       uuid.UUID `gorm:"index"`
	RestaurantID uuid.UUID `gorm:"index"`
}

type RatingData struct {
	Db *gorm.DB
}

func NewRatingData(Db *gorm.DB) *RatingData {
	return &RatingData{Db: Db}
}

// Create rating, inserts into database.
func (ratingData *RatingData) CreateRating(rating Rating) (*Rating, error) {
	result := ratingData.Db.Create(&rating)
	if result.Error != nil {
		return nil, result.Error
	}
	return &rating, nil
}

// Delete ratings.
func (ratingData *RatingData) DeleteRatings(restaurantID uuid.UUID, dishID uuid.UUID) error {
	result := ratingData.Db.Where("dish_id = ? AND restaurant_id = ?", dishID, restaurantID).Delete(&Rating{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
