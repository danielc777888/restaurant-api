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
func (service *RestaurantService) ListRestaurants() ([]RestaurantResult, error) {
	restaurants, err := service.Data.ListRestaurants()
	if err != nil {
		return nil, err
	}

	// map to result array
	result := make([]RestaurantResult, len(restaurants))
	for i, restaurant := range restaurants {
		result[i] = RestaurantResult{
			ID:   restaurant.ID,
			Name: restaurant.Name,
		}
	}
	return result, nil
}
