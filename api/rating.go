package api

import (
	"middleearth/eateries/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createRatingRequest struct {
	Description string `json:"description" binding:"required,min=3,max=200"`
	DishID      string `json:"dishID" binding:"required"`
}

type ratingResponse struct {
	ID           string `json:"id"`
	Description  string `json:"description"`
	DishID       string `json:"dishID"`
	RestaurantID string `json:"restaurantID"`
	Sentiment    string `json:"sentiment"`
}

type RatingAPI struct {
	Service *service.RatingService
}

func NewRatingAPI(Service *service.RatingService) *RatingAPI {
	return &RatingAPI{Service: Service}
}

// CreateRating godoc
// @Summary      Create a rating
// @Description  create a rating
// @Tags         ratings
// @Accept       json
// @Produce      json
// @Param		 RestaurantID	header		string	true	"RestaurantID header"
// @Param		 rating body    api.createRatingRequest   true  "Create rating"
// @Success      200  {array}   api.ratingResponse
// @Router       /ratings [post]
func (api *RatingAPI) CreateRating(ginContext *gin.Context) {

	restaurantID, err := GetRestaurantHeader(ginContext)
	if err != nil {
		ErrorResponse(err, ginContext)
		return
	}

	var request createRatingRequest
	if err := ginContext.BindJSON(&request); err != nil {
		ErrorResponse(err, ginContext)
		return
	}

	// map to rating action
	dishID, _ := uuid.Parse(request.DishID)
	action := service.CreateRatingAction{
		Description: request.Description,
		DishID:      dishID,
	}

	createdRating, err := api.Service.CreateRating(restaurantID, action)
	if err != nil {
		ErrorResponse(err, ginContext)
		return
	}

	// map to rating response
	sentiment := ""
	if createdRating.Sentiment != nil {
		sentiment = *createdRating.Sentiment
	}
	result := ratingResponse{
		ID:           createdRating.ID.String(),
		Description:  createdRating.Description,
		DishID:       createdRating.DishID.String(),
		RestaurantID: createdRating.RestaurantID.String(),
		Sentiment:    sentiment,
	}
	ginContext.IndentedJSON(http.StatusOK, result)
}
