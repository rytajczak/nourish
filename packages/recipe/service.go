package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Service interface {
	SearchRecipes(string, context.Context) []Recipe
	GetRecipeById(context.Context)
	CreateCustomRecipe(context.Context)
}

type RecipeService struct {
	url    string
	client *http.Client
}

func (r *RecipeService) SearchRecipes(query string, ctx context.Context) []Recipe {
	url := fmt.Sprintf("https://spoonacular-recipe-food-nutrition-v1.p.rapidapi.com/recipes/complexSearch?query=%s&instructionsRequired=false&fillIngredients=false&addRecipeInformation=false&addRecipeInstructions=false&addRecipeNutrition=false&ignorePantry=true&offset=0&number=10", query)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", "54749c10eamsh09c102a879618cdp1a0440jsn8aa999a1f52c")
	req.Header.Add("x-rapidapi-host", "spoonacular-recipe-food-nutrition-v1.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer res.Body.Close()

	var response Response
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Fatal("couldn't unmarshal results")
	}

	return response.Recipes
}

func (r *RecipeService) GetRecipeById(context.Context) {
	panic("unimplemented")
}

func (r *RecipeService) CreateCustomRecipe(context.Context) {
	panic("unimplemented")
}

func NewRecipeService(url string) Service {
	return &RecipeService{
		url:    url,
		client: &http.Client{Timeout: 10 * time.Second},
	}
}
