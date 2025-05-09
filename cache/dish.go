package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"middleearth/eateries/data"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type DishCache struct {
	Redis *redis.Client
	Ctx   *context.Context
}

func NewDishCache(Redis *redis.Client, Ctx *context.Context) *DishCache {
	return &DishCache{Redis: Redis, Ctx: Ctx}
}

// Gets list dish cache key
func getListDishKey(restaurantID uuid.UUID) string {
	return fmt.Sprintf("restaurant_dishes:%s", restaurantID.String())
}

// Adds dishes to cache using restaurantID as key
func (dishCache *DishCache) AddDishes(restaurantID uuid.UUID, dishes []data.Dish) {
	key := getListDishKey(restaurantID)
	fmt.Println("SET key to cache:", key)
	dishesJson, _ := json.Marshal(dishes)
	err := dishCache.Redis.Set(*dishCache.Ctx, key, string(dishesJson), time.Second*30).Err()
	if err != nil {
		fmt.Println("Error adding to cache:", err)
	}
}

// Deletes dishes from cache using restaurantID as key
func (dishCache *DishCache) DeleteDishes(restaurantID uuid.UUID) {
	key := getListDishKey(restaurantID)
	fmt.Println("DEL key from cache:", key)
	err := dishCache.Redis.Del(*dishCache.Ctx, key).Err()
	if err != nil {
		fmt.Println("Error deleting from cache:", err)
	}
}

// Get dishes from cache using restaurantID as key
func (dishCache *DishCache) GetDishes(restaurantID uuid.UUID) ([]data.Dish, error) {
	var dishes []data.Dish
	key := getListDishKey(restaurantID)
	fmt.Println("GET key from cache:", key)
	cachedDishes, getError := dishCache.Redis.Get(*dishCache.Ctx, key).Result()
	if getError != nil {
		return dishes, getError
	}
	b := []byte(cachedDishes)
	unmarshalError := json.Unmarshal(b, &dishes)
	if unmarshalError != nil {
		return dishes, unmarshalError
	}
	fmt.Println("GOT DISHES FROM CACHE::", dishes)
	return dishes, nil
}
