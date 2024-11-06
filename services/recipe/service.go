package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Service interface {
	SearchRecipes(string, context.Context) map[string]any
	GetRecipeById(int, context.Context)
	CreateCustomRecipe(context.Context)
}

type RecipeService struct {
	url    string
	host   string
	key    string
	client *http.Client
}

func (r *RecipeService) newRequest(method, endpoint string, body io.Reader) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", r.url, endpoint)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-rapidapi-key", r.key)
	req.Header.Add("x-rapidapi-host", r.host)
	return req, nil
}

func (r *RecipeService) SearchRecipes(query string, ctx context.Context) map[string]any {
	req, err := r.newRequest("GET", "/recipes/complexSearch", nil)
	if err != nil {
		log.Fatal("failed to attach headers")
	}

	q := req.URL.Query()
	q.Add("query", query)
	q.Add("addRecipeNutrition", "true")
	q.Add("instructionsRequired", "true")
	q.Add("number", "30")
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer res.Body.Close()

	var response map[string]any
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Fatal("couldn't unmarshal results")
	}

	return response
}

func (r *RecipeService) GetRecipeById(id int, ctx context.Context) {
	panic("unimplemented")
}

func (r *RecipeService) CreateCustomRecipe(context.Context) {
	panic("unimplemented")
}

func NewRecipeService(host string, key string) Service {
	url := fmt.Sprintf("https://%s", host)

	return &RecipeService{
		url:    url,
		host:   host,
		key:    key,
		client: &http.Client{Timeout: 10 * time.Second},
	}
}
