package main

import "context"

type Service interface {
	SearchRecipesByIngredient(context.Context)
	CreateNewRecipe(context.Context)
	UpdateRecipe(context.Context)
}

type RecipeService struct{}

func NewRecipeService() Service {
	return &RecipeService{}
}

func (s *RecipeService) SearchRecipesByIngredient(context.Context) {

}

func (s *RecipeService) CreateNewRecipe(context.Context) {

}

func (s *RecipeService) UpdateRecipe(context.Context) {

}
