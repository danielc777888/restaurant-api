package api

import (
	"fmt"
	"middleearth/eateries/data"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Dish struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Price        uint   `json:"price"`
	RestaurantID uint   `json:"restaurantID"`
}

type DishAPI struct {
	Db *gorm.DB
}

func NewDishAPI(Db *gorm.DB) *DishAPI {
	return &DishAPI{Db: Db}
}

// @BasePath /api/v1

func (dishApi *DishAPI) CreateDish(c *gin.Context) {
	var newDish Dish
	if err := c.BindJSON(&newDish); err != nil {
		return
	}
	dbDish := mapDishToDB(newDish)
	result := dishApi.Db.Create(&dbDish)
	fmt.Printf("DB result error %s, rows %d", result.Error, result.RowsAffected)
	c.IndentedJSON(http.StatusCreated, newDish)
}

func mapDishToDB(dish Dish) data.Dish {
	return data.Dish{
		Name:         dish.Name,
		Description:  dish.Description,
		Price:        dish.Price,
		RestaurantID: dish.RestaurantID,
	}
}
