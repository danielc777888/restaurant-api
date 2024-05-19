package api

import (
	"errors"
	"fmt"
	"middleearth/eateries/cache"
	"middleearth/eateries/data"
	"middleearth/eateries/env"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Dish struct {
	Name         string `json:"name" binding:"required,min=3,max=20"`
	Description  string `json:"description" binding:"required,min=3,max=200"`
	Price        uint   `json:"price" binding:"required"`
	RestaurantID string `json:"restaurantID" binding:"required"`
}

type UpdateDish struct {
	ID           string `json:"id" binding:"required,min=3,max=20"`
	Name         string `json:"name" binding:"required,min=3,max=200"`
	Description  string `json:"description" binding:"required"`
	Price        uint   `json:"price" binding:"required"`
	RestaurantID string `json:"restaurantID" binding:"required"`
}

type DishAPI struct {
	Db    *gorm.DB
	Cache *cache.DishCache
}

func NewDishAPI(Db *gorm.DB, Cache *cache.DishCache) *DishAPI {
	return &DishAPI{Db: Db, Cache: Cache}
}

// @BasePath /api/v1

func (dishApi *DishAPI) CreateDish(c *gin.Context) {
	restaurantID, err := GetRestaurantHeader(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	var newDish Dish
	if err := c.BindJSON(&newDish); err != nil {
		fmt.Println("Validation error: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "VALIDATION ERROR",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}
	dbDish := mapDishToDB(newDish)
	result := dishApi.Db.Create(&dbDish)
	fmt.Println("DB result error:", result.Error)
	if env.CacheEnabled() {
		dishApi.Cache.DeleteDishes(restaurantID)
	}
	c.IndentedJSON(http.StatusOK, newDish)
}

func (dishApi *DishAPI) UpdateDish(c *gin.Context) {
	restaurantID, err := GetRestaurantHeader(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	var dish UpdateDish
	if err := c.BindJSON(&dish); err != nil {
		fmt.Printf("Bind error %s", err)
		return
	}
	dbDish := mapUpdateDishToDB(dish)
	result := dishApi.Db.Where("id = ? AND restaurant_id = ?", dbDish.ID, restaurantID).Save(&dbDish)
	fmt.Printf("DB result error %s, rows %d", result.Error, result.RowsAffected)
	if env.CacheEnabled() {
		dishApi.Cache.DeleteDishes(restaurantID)
	}
	c.IndentedJSON(http.StatusOK, dbDish)
}

func (dishApi *DishAPI) DeleteDish(c *gin.Context) {
	restaurantID, err := GetRestaurantHeader(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	id, _ := uuid.Parse(c.Param("id"))
	var dish data.Dish
	result := dishApi.Db.Where("id = ? AND restaurant_id = ?", id, restaurantID).First(&dish)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Dish not found"})
		return
	}

	// TODO: Use db transactions

	// delete associated ratings first
	dishApi.Db.Where("dish_id = ?", dish.ID).Delete(&data.Rating{})

	deleteResult := dishApi.Db.Delete(&dish)
	if deleteResult.Error != nil {
		fmt.Println("ERROR: Deleting Dish: ", deleteResult.Error)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error deleting dish"})
		return
	}
	if env.CacheEnabled() {
		dishApi.Cache.DeleteDishes(restaurantID)
	}
	c.IndentedJSON(http.StatusOK, id)
}

func (dishApi *DishAPI) GetDish(c *gin.Context) {
	restaurantID, err := GetRestaurantHeader(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	id := c.Param("id")
	var dish data.Dish
	result := dishApi.Db.Where("id = ? AND restaurant_id = ?", id, restaurantID).First(&dish)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Dish not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, dish)
}

func (dishApi *DishAPI) ListDish(c *gin.Context) {
	restaurantID, err := GetRestaurantHeader(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	var dishes []data.Dish
	if env.CacheEnabled() {
		dishes, err := dishApi.Cache.GetDishes(restaurantID)
		if dishes == nil {
			fmt.Println("Getting dishes from DATABASE:", err)
			dishApi.Db.Where("restaurant_id = ?", restaurantID).Find(&dishes)
			dishApi.Cache.AddDishes(restaurantID, dishes)
		}
		c.IndentedJSON(http.StatusOK, mapDishesToJSON(dishes))
		return
	}
	dishApi.Db.Find(&dishes)
	c.IndentedJSON(http.StatusOK, mapDishesToJSON(dishes))
}

func mapDishToDB(dish Dish) data.Dish {
	restaurantID, _ := uuid.Parse(dish.RestaurantID)
	return data.Dish{
		Name:         dish.Name,
		Description:  dish.Description,
		Price:        dish.Price,
		RestaurantID: restaurantID,
	}
}

func mapUpdateDishToDB(dish UpdateDish) data.Dish {
	restaurantID, _ := uuid.Parse(dish.RestaurantID)
	dishID, _ := uuid.Parse(dish.ID)
	return data.Dish{
		ID:           dishID,
		Name:         dish.Name,
		Description:  dish.Description,
		Price:        dish.Price,
		RestaurantID: restaurantID,
	}
}

func mapDishesToJSON(dishes []data.Dish) []Dish {
	var result []Dish
	for _, dish := range dishes {
		result = append(result, mapDishToJSON(dish))
	}
	return result
}

func mapDishToJSON(dish data.Dish) Dish {
	return Dish{
		Name:         dish.Name,
		Description:  dish.Description,
		Price:        dish.Price,
		RestaurantID: dish.RestaurantID.String(),
	}
}
