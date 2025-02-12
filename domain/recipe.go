package domain

import "cookieBot/bootstrap"

type Recipe struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Image     string `json:"image"`
	ImageType string `json:"imageType"`
}

type RecipeUseCase interface {
	GetRecipesByType(recipeName string, config bootstrap.Config) ([]Recipe, error)
}
