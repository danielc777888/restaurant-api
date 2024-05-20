package api

import (
	"errors"
	"middleearth/eateries/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createDishRequest struct {
	Name         string `json:"name" binding:"required,min=3,max=20"`
	Description  string `json:"description" binding:"required,min=3,max=200"`
	Price        uint   `json:"price" binding:"required"`
	RestaurantID string `json:"restaurantID" binding:"required"`
}

type updateDishRequest struct {
	ID           string `json:"id" binding:"required"`
	Name         string `json:"name" binding:"required,min=3,max=50"`
	Description  string `json:"description" binding:"required,min=3,max=200"`
	Price        uint   `json:"price" binding:"required"`
	RestaurantID string `json:"restaurantID" binding:"required"`
}

type dishResponse struct {
	ID           string
	Name         string
	Description  string
	Price        uint
	RestaurantID string
}

type DishAPI struct {
	Service *service.DishService
}

func NewDishAPI(Service *service.DishService) *DishAPI {
	return &DishAPI{Service: Service}
}

// CreateDish godoc
// @Summary      Create a dish
// @Description  create a dish
// @Tags         dishes
// @Accept       json
// @Produce      json
// @Param		 Authorization	header		string	true	"Authentication header"
// @Param		 RestaurantID	header		string	true	"RestaurantID header"
// @Param		 dish body    api.createDishRequest   true  "Create dish"
// @Success      200  {object}   api.dishResponse
// @Router       /dishes [post]
func (dishApi *DishAPI) CreateDish(ginContext *gin.Context) {
	restaurantID, err := GetRestaurantHeader(ginContext)
	if err != nil {
		ErrorResponse(err, ginContext)
		return
	}
	var request createDishRequest
	if err := ginContext.BindJSON(&request); err != nil {
		ErrorResponse(err, ginContext)
		return
	}

	jsonID, _ := uuid.Parse(request.RestaurantID)
	validationErr := validateRestaurantID(restaurantID, jsonID)
	if validationErr != nil {
		ErrorResponse(validationErr, ginContext)
		return
	}

	// map to dish action
	action := service.CreateDishAction{
		Name:         request.Name,
		Description:  request.Description,
		Price:        request.Price,
		RestaurantID: restaurantID,
	}

	result, err := dishApi.Service.CreateDish(restaurantID, action)
	if err != nil {
		ErrorResponse(err, ginContext)
		return
	}

	response := mapToDishResponse(*result)

	ginContext.IndentedJSON(http.StatusOK, response)
}

// UpdateDish godoc
// @Summary      Update a dish
// @Description  updated a dish
// @Tags         dishes
// @Accept       json
// @Produce      json
// @Param		 Authorization	header		string	true	"Authentication header"
// @Param		 RestaurantID	header		string	true	"RestaurantID header"
// @Param		 dish body    api.updateDishRequest   true  "Update dish"
// @Success      200  {object}   api.dishResponse
// @Router       /dishes [patch]
func (dishApi *DishAPI) UpdateDish(ginContext *gin.Context) {

	restaurantID, err := GetRestaurantHeader(ginContext)
	if err != nil {
		ErrorResponse(err, ginContext)
		return
	}
	var request updateDishRequest
	if err := ginContext.BindJSON(&request); err != nil {
		ErrorResponse(err, ginContext)
		return
	}

	jsonID, _ := uuid.Parse(request.RestaurantID)
	validationErr := validateRestaurantID(restaurantID, jsonID)
	if validationErr != nil {
		ErrorResponse(validationErr, ginContext)
		return
	}

	// map to dish action
	dishID, _ := uuid.Parse(request.ID)
	action := service.UpdateDishAction{
		ID:           dishID,
		Name:         request.Name,
		Description:  request.Description,
		Price:        request.Price,
		RestaurantID: restaurantID,
	}

	result, err := dishApi.Service.UpdateDish(restaurantID, action)
	if err != nil {
		ErrorResponse(err, ginContext)
		return
	}

	response := mapToDishResponse(*result)

	ginContext.IndentedJSON(http.StatusOK, response)
}

// DeleteDish godoc
// @Summary      Delete a dish
// @Description  delete a dish
// @Tags         dishes
// @Accept       json
// @Produce      json
// @Param		 Authorization	header		string	true	"Authentication header"
// @Param		 RestaurantID	header		string	true	"RestaurantID header"
// @Param		 dish_id	path		string		true	"Dish ID"
// @Success      200
// @Router       /dishes/{dish_id} [delete]
func (dishApi *DishAPI) DeleteDish(ginContext *gin.Context) {
	restaurantID, err := GetRestaurantHeader(ginContext)
	if err != nil {
		ErrorResponse(err, ginContext)
		return
	}
	id := ginContext.Param("id")
	dishID, _ := uuid.Parse(id)

	deleteError := dishApi.Service.DeleteDish(restaurantID, dishID)
	if deleteError != nil {
		ErrorResponse(deleteError, ginContext)
		return
	}

	ginContext.IndentedJSON(http.StatusOK, gin.H{
		"id": id,
	})
}

// GetDish godoc
// @Summary      Get a dish
// @Description  get a dish
// @Tags         dishes
// @Accept       json
// @Produce      json
// @Param		 RestaurantID	header		string	true	"RestaurantID header"
// @Param		 dish_id	path		string		true	"Dish ID"
// @Success      200  {object}   api.dishResponse
// @Router       /dishes/{dish_id} [get]
func (dishApi *DishAPI) GetDish(ginContext *gin.Context) {
	restaurantID, err := GetRestaurantHeader(ginContext)
	if err != nil {
		ErrorResponse(err, ginContext)
		return
	}
	id := ginContext.Param("id")
	dishID, _ := uuid.Parse(id)
	retrievedDish, err := dishApi.Service.GetDish(restaurantID, dishID)
	if err != nil {
		ErrorResponse(err, ginContext)
		return
	}
	response := mapToDishResponse(*retrievedDish)
	ginContext.IndentedJSON(http.StatusOK, response)
}

// ListDishes godoc
// @Summary      List dishes
// @Description  list dishes
// @Tags         dishes
// @Accept       json
// @Produce      json
// @Param		 RestaurantID	header		string	true	"RestaurantID header"
// @Success      200  {array}   api.dishResponse
// @Router       /dishes [get]
func (dishApi *DishAPI) ListDishes(ginContext *gin.Context) {
	restaurantID, err := GetRestaurantHeader(ginContext)
	if err != nil {
		ErrorResponse(err, ginContext)
		return
	}
	dishes, err := dishApi.Service.ListDishes(restaurantID)
	if err != nil {
		ErrorResponse(err, ginContext)
		return
	}

	// map []service.DishResult to []api.dishResponse
	var response []dishResponse
	for _, dish := range dishes {
		response = append(response, mapToDishResponse(dish))
	}
	ginContext.IndentedJSON(http.StatusOK, response)
}

// Maps service.DishResult to api.dishResponse
func mapToDishResponse(result service.DishResult) dishResponse {
	return dishResponse{
		ID:           result.ID.String(),
		Name:         result.Name,
		Description:  result.Description,
		Price:        result.Price,
		RestaurantID: result.RestaurantID.String(),
	}
}

// Validates the restaurant ids used in header and json payload.
// Checks that they are equal and returns an error
func validateRestaurantID(headerID uuid.UUID, jsonID uuid.UUID) error {
	if headerID != jsonID {
		return errors.New("invalid restaurant id")
	}
	return nil
}
