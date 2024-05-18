package api

import (
	"fmt"
	"middleearth/eateries/data"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateRating struct {
	Description string `json:"description"`
	DishID      string `json:"dishID"`
}

type Rating struct {
	ID           string `json:"id"`
	Description  string `json:"description"`
	DishID       string `json:"dishID"`
	RestaurantID string `json:"restaurantID"`
}

type RatingAPI struct {
	Db *gorm.DB
}

func NewRatingAPI(Db *gorm.DB) *RatingAPI {
	return &RatingAPI{Db: Db}
}

// @BasePath /api/v1

func (ratingApi *RatingAPI) CreateRating(c *gin.Context) {
	restaurantID, err := GetRestaurantHeader(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	var rating CreateRating
	if err := c.BindJSON(&rating); err != nil {
		return
	}
	dbRating := mapCreateRatingToDB(restaurantID, rating)
	result := ratingApi.Db.Create(&dbRating)
	fmt.Printf("DB result error %s, rows %d", result.Error, result.RowsAffected)
	c.IndentedJSON(http.StatusOK, mapRatingToJSON(dbRating))
}

func mapCreateRatingToDB(restaurantID uuid.UUID, rating CreateRating) data.Rating {
	dishID, _ := uuid.Parse(rating.DishID)
	return data.Rating{
		ID:           uuid.New(),
		Description:  rating.Description,
		DishID:       dishID,
		RestaurantID: restaurantID,
	}
}

func mapRatingToJSON(rating data.Rating) Rating {
	return Rating{
		ID:          rating.ID.String(),
		Description: rating.Description,
		DishID:      rating.DishID.String(),
	}
}
