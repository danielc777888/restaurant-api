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
