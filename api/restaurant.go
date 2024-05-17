package api

import (
	"fmt"
	"net/http"

	"middleearth/eateries/data"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Restaurant struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type RestaurantAPI struct {
	Db *gorm.DB
}

func NewRestaurantAPI(Db *gorm.DB) *RestaurantAPI {
	return &RestaurantAPI{Db: Db}
}

// var restaurants = []Restaurant{
// 	{ID: "1", Name: "Rest1"},
// 	{ID: "2", Name: "Rest2"},
// }

// var dsn = "host=localhost user=dancingponysvc password=password dbname=dancingpony port=5432"
// var db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

// @BasePath /api/v1

// GetRestaurants godoc
// @Summary Gets a list of restaurants
// @Schemes
// @Description get restaurants
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {Restaurant} []Restaurant
// @Router /restaurants [get]
func (r *RestaurantAPI) GetRestaurants(c *gin.Context) {
	fmt.Println("Getting restaurants")
	var restaurants []data.Restaurant
	r.Db.Find(&restaurants)
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
