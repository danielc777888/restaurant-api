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
func GetRestaurantHeader(ginContext *gin.Context) (uuid.UUID, error) {
	header := ginContext.Request.Header.Get("RestaurantID")
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

// ListRestaurants godoc
// @Summary      List restaurants
// @Description  list restaurants
// @Tags         restaurants
// @Accept       json
// @Produce      json
// @Success      200  {array}   api.restaurantResponse
// @Router       /restaurants [get]
func (api *RestaurantAPI) ListRestaurants(ginContext *gin.Context) {
	restaurants, err := api.Service.ListRestaurants()
	if err != nil {
		ErrorResponse(err, ginContext)
		return
	}

	// map to response
	response := make([]restaurantResponse, len(restaurants))
	for i, restaurant := range restaurants {
		response[i] = restaurantResponse{
			ID:   restaurant.ID.String(),
			Name: restaurant.Name,
		}
	}
	ginContext.IndentedJSON(http.StatusOK, response)
}
