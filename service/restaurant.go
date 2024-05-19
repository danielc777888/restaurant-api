package service

import (
	"middleearth/eateries/data"

	"github.com/google/uuid"
)

type RestaurantResult struct {
	ID   uuid.UUID
	Name string
}

type RestaurantService struct {
	Data *data.RestaurantData
}

func NewRestaurantService(Data *data.RestaurantData) *RestaurantService {
	return &RestaurantService{Data: Data}
}

// List restaurants
func (r *RestaurantService) ListRestaurants() ([]RestaurantResult, error) {
	result, err := r.Data.ListRestaurants()
	if err != nil {
		return nil, err
	}
	return mapToResult(result), nil
}

// Maps data.Restaurant array to service.RestaurantResult array
func mapToResult(restaurants []data.Restaurant) []RestaurantResult {
	result := make([]RestaurantResult, len(restaurants))
	for i, restaurant := range restaurants {
		result[i] = RestaurantResult{
			ID:   restaurant.ID,
			Name: restaurant.Name,
		}
	}
	return result
}
