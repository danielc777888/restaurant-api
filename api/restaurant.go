package api

import (
	"errors"
	"fmt"
	"net/http"

	"middleearth/eateries/data"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Restaurant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RestaurantAPI struct {
	Db *gorm.DB
}

func NewRestaurantAPI(Db *gorm.DB) *RestaurantAPI {
	return &RestaurantAPI{Db: Db}
}

func GetRestaurantHeader(c *gin.Context) (uuid.UUID, error) {
	header := c.Request.Header.Get("RestaurantID")
	if header == "" {
		fmt.Println("Invalid restaurant header")
		return uuid.Nil, errors.New("restaurant header not found")
	}
	restaurantID, err := uuid.Parse(header)
	if err != nil {
		fmt.Println("Invalid restaurant header")
		return uuid.Nil, errors.New("invalid restaurant header")
	}
	return restaurantID, nil
}

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
			ID:   restaurant.ID.String(),
			Name: restaurant.Name,
		}
	}
	return result
}
