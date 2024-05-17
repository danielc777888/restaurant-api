package api

import (
	"fmt"
	"middleearth/eateries/data"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateRating struct {
	Description string `json:"description"`
	DishID      uint   `json:"dishID"`
}

type Rating struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	DishID      uint   `json:"dishID"`
}

type RatingAPI struct {
	Db *gorm.DB
}

func NewRatingAPI(Db *gorm.DB) *RatingAPI {
	return &RatingAPI{Db: Db}
}

// @BasePath /api/v1

func (ratingApi *RatingAPI) CreateRating(c *gin.Context) {
	var rating CreateRating
	if err := c.BindJSON(&rating); err != nil {
		return
	}
	dbRating := mapCreateRatingToDB(rating)
	result := ratingApi.Db.Create(&dbRating)
	fmt.Printf("DB result error %s, rows %d", result.Error, result.RowsAffected)
	c.IndentedJSON(http.StatusOK, mapRatingToJSON(dbRating))
}

// func (dishApi *DishAPI) UpdateDish(c *gin.Context) {
// 	var dish UpdateDish
// 	if err := c.BindJSON(&dish); err != nil {
// 		fmt.Printf("Bind error %s", err)
// 		return
// 	}
// 	dbDish := mapUpdateDishToDB(dish)
// 	result := dishApi.Db.Save(&dbDish)
// 	fmt.Printf("DB result error %s, rows %d", result.Error, result.RowsAffected)
// 	c.IndentedJSON(http.StatusOK, dbDish)
// }

// func (dishApi *DishAPI) DeleteDish(c *gin.Context) {
// 	id := c.Param("id")
// 	var dish data.Dish
// 	result := dishApi.Db.First(&dish, id)
// 	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Dish not found"})
// 		return
// 	}
// 	dishApi.Db.Delete(&dish)
// 	c.IndentedJSON(http.StatusOK, id)
// }

// func (dishApi *DishAPI) GetDish(c *gin.Context) {
// 	id := c.Param("id")
// 	var dish data.Dish
// 	result := dishApi.Db.First(&dish, id)
// 	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Dish not found"})
// 		return
// 	}
// 	c.IndentedJSON(http.StatusOK, dish)
// }

// func (dishApi *DishAPI) ListDish(c *gin.Context) {
// 	var dishes []data.Dish
// 	dishApi.Db.Find(&dishes)
// 	c.IndentedJSON(http.StatusOK, mapDishesToJSON(dishes))
// }

func mapCreateRatingToDB(rating CreateRating) data.Rating {
	return data.Rating{
		Description: rating.Description,
		DishID:      rating.DishID,
	}
}

func mapRatingToJSON(rating data.Rating) Rating {
	return Rating{
		ID:          rating.ID,
		Description: rating.Description,
		DishID:      rating.DishID,
	}
}

// func mapUpdateDishToDB(dish UpdateDish) data.Dish {
// 	return data.Dish{
// 		ID:           dish.ID,
// 		Name:         dish.Name,
// 		Description:  dish.Description,
// 		Price:        dish.Price,
// 		RestaurantID: dish.RestaurantID,
// 	}
// }

// func mapDishesToJSON(dishes []data.Dish) []Dish {
// 	var result []Dish
// 	for _, dish := range dishes {
// 		result = append(result, mapDishToJSON(dish))
// 	}
// 	return result
// }

// func mapDishToJSON(dish data.Dish) Dish {
// 	return Dish{
// 		Name:         dish.Name,
// 		Description:  dish.Description,
// 		Price:        dish.Price,
// 		RestaurantID: dish.RestaurantID,
// 	}
// }
