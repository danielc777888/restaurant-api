package service

import (
	"middleearth/eateries/data"
	"middleearth/eateries/env"
	"middleearth/eateries/llm"

	"github.com/google/uuid"
)

type CreateRatingAction struct {
	Description string
	DishID      uuid.UUID
}

type RatingResult struct {
	ID           uuid.UUID
	Description  string
	DishID       uuid.UUID
	RestaurantID uuid.UUID
	Sentiment    *string
}

type RatingService struct {
	Data *data.RatingData
}

func NewRatingService(Data *data.RatingData) *RatingService {
	return &RatingService{Data: Data}
}

// Create a rating from an action in a specific restaurant
// Returns a rating result or error.
func (service *RatingService) CreateRating(restaurantID uuid.UUID, action CreateRatingAction) (*RatingResult, error) {
	// map to rating data
	rating := data.Rating{
		ID:           uuid.New(),
		RestaurantID: restaurantID,
		DishID:       action.DishID,
		Description:  action.Description,
	}

	// sentiment analysis using a LLM (large language model)
	if env.LLMEnabled() {
		sentiment, err := llm.AnalyzeSentiment(rating.Description)
		if err != nil {
			return nil, err
		}
		rating.Sentiment = sentiment
	}

	createdRating, err := service.Data.CreateRating(rating)
	if err != nil {
		return nil, err
	}

	// map to rating result
	result := RatingResult{
		ID:           createdRating.ID,
		Description:  createdRating.Description,
		DishID:       createdRating.DishID,
		RestaurantID: createdRating.RestaurantID,
		Sentiment:    createdRating.Sentiment,
	}
	return &result, nil
}
