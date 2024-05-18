package api

import (
	"errors"
	"fmt"
	"middleearth/eateries/cache"
	"middleearth/eateries/data"
	"middleearth/eateries/env"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Dish struct {
	Name         string `json:"name" binding:"required,min=3,max=20"`
	Description  string `json:"description" binding:"required,min=3,max=200"`
	Price        uint   `json:"price" binding:"required"`
	RestaurantID uint   `json:"restaurantID" binding:"required"`
}

type UpdateDish struct {
	ID           uint   `json:"id" binding:"required,min=3,max=20"`
	Name         string `json:"name" binding:"required,min=3,max=200"`
	Description  string `json:"description" binding:"required"`
	Price        uint   `json:"price" binding:"required"`
	RestaurantID uint   `json:"restaurantID" binding:"required"`
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
		dishApi.Cache.DeleteDishes(1)
	}
	c.IndentedJSON(http.StatusOK, newDish)
}

func (dishApi *DishAPI) UpdateDish(c *gin.Context) {
	var dish UpdateDish
	if err := c.BindJSON(&dish); err != nil {
		fmt.Printf("Bind error %s", err)
		return
	}
	dbDish := mapUpdateDishToDB(dish)
	result := dishApi.Db.Save(&dbDish)
	fmt.Printf("DB result error %s, rows %d", result.Error, result.RowsAffected)
	if env.CacheEnabled() {
		dishApi.Cache.DeleteDishes(1)
	}
	c.IndentedJSON(http.StatusOK, dbDish)
}

func (dishApi *DishAPI) DeleteDish(c *gin.Context) {
	id := c.Param("id")
	var dish data.Dish
	result := dishApi.Db.First(&dish, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Dish not found"})
		return
	}
	dishApi.Db.Delete(&dish)
	if env.CacheEnabled() {
		dishApi.Cache.DeleteDishes(1)
	}
	c.IndentedJSON(http.StatusOK, id)
}

func (dishApi *DishAPI) GetDish(c *gin.Context) {
	id := c.Param("id")
	var dish data.Dish
	result := dishApi.Db.First(&dish, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Dish not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, dish)
}

func (dishApi *DishAPI) ListDish(c *gin.Context) {
	var dishes []data.Dish
	if env.CacheEnabled() {
		dishes, err := dishApi.Cache.GetDishes(1)
		if dishes == nil {
			fmt.Println("Getting dishes from DATABASE:", err)
			dishApi.Db.Find(&dishes)
			dishApi.Cache.AddDishes(1, dishes)
		}
		c.IndentedJSON(http.StatusOK, mapDishesToJSON(dishes))
		return
	}
	dishApi.Db.Find(&dishes)
	c.IndentedJSON(http.StatusOK, mapDishesToJSON(dishes))
}

func mapDishToDB(dish Dish) data.Dish {
	return data.Dish{
		Name:         dish.Name,
		Description:  dish.Description,
		Price:        dish.Price,
		RestaurantID: dish.RestaurantID,
	}
}

func mapUpdateDishToDB(dish UpdateDish) data.Dish {
	return data.Dish{
		ID:           dish.ID,
		Name:         dish.Name,
		Description:  dish.Description,
		Price:        dish.Price,
		RestaurantID: dish.RestaurantID,
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
		RestaurantID: dish.RestaurantID,
	}
}
