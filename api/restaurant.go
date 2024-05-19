package api

import (
	"errors"
	"fmt"
	"net/http"

	"middleearth/eateries/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type restaurantResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RestaurantAPI struct {
	Service *service.RestaurantService
}

func NewRestaurantAPI(Service *service.RestaurantService) *RestaurantAPI {
	return &RestaurantAPI{Service: Service}
}

// Get RestaurantID from gin context request header
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

// LisRestaurants godoc
// @Summary      List restaurants
// @Description  list restaurants
// @Tags         restaurants
// @Accept       json
// @Produce      json
// @Success      200  {array}   api.restaurantResponse
// @Router       /restaurants [get]
func (r *RestaurantAPI) ListRestaurants(c *gin.Context) {
	result, err := r.Service.ListRestaurants()
	if err != nil {
		ErrorResponse(err, c)
		return
	}
	c.IndentedJSON(http.StatusOK, mapToResponse(result))
}

// Map service.RestaurantResult array to api.restaurantResponse array
func mapToResponse(restaurants []service.RestaurantResult) []restaurantResponse {
	result := make([]restaurantResponse, len(restaurants))
	for i, restaurant := range restaurants {
		result[i] = restaurantResponse{
			ID:   restaurant.ID.String(),
			Name: restaurant.Name,
		}
	}
	return result
}
