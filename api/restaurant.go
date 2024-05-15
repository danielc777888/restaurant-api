package api

import (
	"fmt"
	"net/http"

	"middleearth/eateries/data"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Restaurant struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// var restaurants = []Restaurant{
// 	{ID: "1", Name: "Rest1"},
// 	{ID: "2", Name: "Rest2"},
// }

var dsn = "host=localhost user=dancingponysvc password=password dbname=dancingpony port=5432"
var db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

func GetRestaurants(c *gin.Context) {
	fmt.Println("Getting restaruants")
	var restaurants []data.Restaurant
	if err != nil {
		panic("Failed to connect database")
	}
	db.Find(&restaurants)
	c.IndentedJSON(http.StatusOK, mapRestaurantsToJSON(restaurants))
}

func mapRestaurantsToJSON(restaurants []data.Restaurant) []Restaurant {
	result := make([]Restaurant, len(restaurants))
	for i, restaurant := range restaurants {
		result[i] = Restaurant{
			ID:   restaurant.ID,
			Name: restaurant.Name,
		}
	}
	return result
}
