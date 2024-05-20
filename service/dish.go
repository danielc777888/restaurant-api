package service

import (
	"fmt"
	"middleearth/eateries/cache"
	"middleearth/eateries/data"
	"middleearth/eateries/env"

	"github.com/google/uuid"
)

type CreateDishAction struct {
	Name         string
	Description  string
	Price        uint
	RestaurantID uuid.UUID
}

type UpdateDishAction struct {
	ID           uuid.UUID
	Name         string
	Description  string
	Price        uint
	RestaurantID uuid.UUID
}

type DishResult struct {
	ID           uuid.UUID
	Name         string
	Description  string
	Price        uint
	RestaurantID uuid.UUID
}

type DishService struct {
	Data       *data.DishData
	RatingData *data.RatingData
	Cache      *cache.DishCache
}

func NewDishService(Data *data.DishData, RatingData *data.RatingData, Cache *cache.DishCache) *DishService {
	return &DishService{Data: Data, RatingData: RatingData, Cache: Cache}
}

// Create a dish from an action for a specific restaurant
// Returns a dish result or error.
func (service *DishService) CreateDish(restaurantID uuid.UUID, action CreateDishAction) (*DishResult, error) {
	// map to dish data
	dish := data.Dish{
		ID:           uuid.New(),
		RestaurantID: restaurantID,
		Name:         action.Name,
		Description:  action.Description,
		Price:        action.Price,
	}

	createdDish, err := service.Data.CreateDish(dish)
	if err != nil {
		return nil, err
	}

	// caching
	if env.CacheEnabled() {
		service.Cache.DeleteDishes(restaurantID)
	}

	result := mapToDishResult(*createdDish)
	return &result, nil
}

// Update a dish from an action for a specific restaurant.
// Returns a dish result or error.
func (service *DishService) UpdateDish(restaurantID uuid.UUID, action UpdateDishAction) (*DishResult, error) {

	var dish data.Dish
	retrievedDish, err := service.Data.GetDish(restaurantID, action.ID, dish)
	if err != nil {
		return nil, err
	}
	// map to dish data
	retrievedDish.Name = action.Name
	retrievedDish.Description = action.Description
	retrievedDish.Price = action.Price

	updatedDish, err := service.Data.UpdateDish(restaurantID, *retrievedDish)
	if err != nil {
		return nil, err
	}

	// caching
	if env.CacheEnabled() {
		service.Cache.DeleteDishes(restaurantID)
	}

	result := mapToDishResult(*updatedDish)
	return &result, nil
}

// Delete a dish for a specific restaurant.
func (service *DishService) DeleteDish(restaurantID uuid.UUID, dishID uuid.UUID) error {

	// TODO: Use a db transaction

	// delete associated ratings first
	err := service.RatingData.DeleteRatings(restaurantID, dishID)
	if err != nil {
		return err
	}

	// delete dish
	deleteError := service.Data.DeleteDish(restaurantID, dishID)
	if deleteError != nil {
		return deleteError
	}

	// clear cache
	if env.CacheEnabled() {
		service.Cache.DeleteDishes(restaurantID)
	}
	return nil
}

// Get a dish from an action for a specific restaurant.
// Returns a dish result or error.
func (service *DishService) GetDish(restaurantID uuid.UUID, dishID uuid.UUID) (*DishResult, error) {
	var dish data.Dish
	retrievedDish, err := service.Data.GetDish(restaurantID, dishID, dish)
	if err != nil {
		return nil, err
	}
	result := mapToDishResult(*retrievedDish)
	return &result, nil
}

// List dishes for a specific restaurant.
// Returns a dish result list or error.
func (service *DishService) ListDishes(restaurantID uuid.UUID) ([]DishResult, error) {
	if env.CacheEnabled() {
		cachedDishes, err := service.Cache.GetDishes(restaurantID)
		if cachedDishes == nil {
			fmt.Println("Getting dishes from DATABASE:", err)
			retrievedDishes, err := service.Data.ListDishes(restaurantID)
			if err != nil {
				return nil, err
			}
			service.Cache.AddDishes(restaurantID, retrievedDishes)
			result := mapToDishResults(retrievedDishes)
			return result, nil
		}
		result := mapToDishResults(cachedDishes)
		return result, nil
	}
	retrievedDishes, err := service.Data.ListDishes(restaurantID)
	if err != nil {
		return nil, err
	}
	result := mapToDishResults(retrievedDishes)
	return result, nil
}

// Maps []data.Dish to []service.DishResult
func mapToDishResults(dishes []data.Dish) []DishResult {
	var result []DishResult
	for _, dish := range dishes {
		result = append(result, mapToDishResult(dish))
	}
	return result
}

// Maps data.Dish to service.DishResult
func mapToDishResult(dish data.Dish) DishResult {
	result := DishResult{
		ID:           dish.ID,
		Description:  dish.Description,
		Name:         dish.Name,
		Price:        dish.Price,
		RestaurantID: dish.RestaurantID,
	}
	return result
}
