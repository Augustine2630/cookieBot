package usecase

import (
	"cookieBot/bootstrap"
	"cookieBot/domain/response"
	"fmt"
	"resty.dev/v3"
)

type SpoonClient struct {
	Client *resty.Client
}

func NewSpoonClient() *SpoonClient {
	return &SpoonClient{Client: resty.New()}
}

func (s *SpoonClient) GetRecipesByName(recipeName string, config bootstrap.Config) (*response.SpoonacularRecipeResponse, error) {

	recipeResponse := &response.SpoonacularRecipeResponse{}

	res, err := s.Client.R().
		EnableTrace().
		SetQueryParam("apiKey", config.App.SpoonacularToken).
		SetQueryParam("query", recipeName).
		SetQueryParam("number", "5").
		SetResult(recipeResponse). // Set the result to the recipeResponse struct
		Get("https://api.spoonacular.com/recipes/complexSearch")

	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, fmt.Errorf("API request failed with status code: %d", res.StatusCode())
	}

	// Return the parsed response
	return recipeResponse, nil
}
