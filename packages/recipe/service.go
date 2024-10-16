package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Service interface {
	GetRandomRecipes(context.Context) map[string]interface{}
	SearchRecipes(context.Context) map[string]interface{}
}

type RecipeService struct {
	url    string
	client *http.Client
}

func (r *RecipeService) GetRandomRecipes(context.Context) map[string]interface{} {
	req, _ := http.NewRequest("GET", r.url+"/recipes/random?number=10", nil)
	req.Header.Add("x-rapidapi-key", os.Getenv("RAPIDAPI_KEY"))
	req.Header.Add("x-rapidapi-host", os.Getenv("RAPIDAPI_HOST"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		println(err.Error())
		return nil
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal("failed to read json")
	}

	return result
}

func (r *RecipeService) SearchRecipes(context.Context) map[string]interface{} {
	panic("unimplemented")
}

func NewRecipeService(url string) Service {
	return &RecipeService{
		url:    url,
		client: &http.Client{Timeout: 10 * time.Second},
	}
}
