package llm

import (
	"context"
	"fmt"
	"middleearth/eateries/env"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// Analyzes sentiment of a rating.
// It returns a string of either POSITIVE, NEGATIVE or NEUTRAL.
func AnalyzeSentiment(rating string) (*string, error) {
	ctx := context.Background()
	apKey := env.GeminiAPIKey()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apKey))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")
	prompt := fmt.Sprintf("Analyze the sentiment of the following Restaurant Dish Review and Classify it as POSITIVE, NEGATIVE, or NEUTRAL. '%s'", rating)
	//fmt.Println("Here is the prompt to user:", prompt)
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, err
	}
	//fmt.Println("****GEMINI_PRO***::", resp.Candidates[0].Content.Parts[0])
	classification := fmt.Sprint(resp.Candidates[0].Content.Parts[0])
	return &classification, nil
}
