package api

import (
	"context"
	"fmt"
	"log"
	"middleearth/eateries/data"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"github.com/google/uuid"
	"google.golang.org/api/option"
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

	// TODO: Takes quite long, need to use goroutines??
	// get sentiment from gemini
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// For text-only input, use the gemini-pro model,

	model := client.GenerativeModel("gemini-pro")
	prompt := fmt.Sprintf("Analyze the sentiment of the following Restaurant Dish Review and Classify it as POSITIVE, NEGATIVE, or NEUTRAL. '%s'", dbRating.Description)
	fmt.Println("Here is the prompt to user:", prompt)
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("****GEMINI_PRO***::", resp.Candidates[0].Content.Parts[0])
	classification := fmt.Sprint(resp.Candidates[0].Content.Parts[0])
	dbRating.Sentiment = &classification

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
