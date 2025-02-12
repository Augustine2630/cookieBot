package usecase

import (
	"cookieBot/bootstrap"
	"cookieBot/domain"
)

type RecipeUsecase struct {
}

func (r RecipeUsecase) GetRecipesByType(recipeName string, config bootstrap.Config) ([]domain.Recipe, error) {
	client := NewSpoonClient()

	var recipes = make([]domain.Recipe, 0)

	responseRecipes, err := client.GetRecipesByName(recipeName, config)
	if err != nil {
		return nil, err
	}

	for _, result := range responseRecipes.Results {
		recipes = append(recipes, domain.Recipe{
			Id:        result.Id,
			Title:     result.Title,
			Image:     result.Image,
			ImageType: result.ImageType,
		})
	}

	return recipes, nil
}

func NewRecipeUsecase() domain.RecipeUseCase {
	return &RecipeUsecase{}
}
