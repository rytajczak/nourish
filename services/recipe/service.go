package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Service interface {
	SearchRecipes(params url.Values, ctx context.Context) map[string]any
	GetRecipeInfo(id int, ctx context.Context) (map[string]any, bool, error)
	GetRecipeInfoBulk(ids []int, ctx context.Context) map[string]any
}

type RecipeService struct {
	url   string
	host  string
	key   string
	cache *Cache
}

func NewRecipeService(host string, key string, cache *Cache) Service {
	url := fmt.Sprintf("https://%s", host)

	return &RecipeService{
		url:   url,
		host:  host,
		key:   key,
		cache: cache,
	}
}

func (r *RecipeService) newRequest(method string, endpoint string) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", r.url, endpoint)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-rapidapi-key", r.key)
	req.Header.Add("x-rapidapi-host", r.host)
	return req, nil
}

func (s *RecipeService) SearchRecipes(params url.Values, ctx context.Context) map[string]any {
	req, err := s.newRequest("GET", "/recipes/complexSearch")
	if err != nil {
		log.Fatal("failed to attach headers")
	}

	q := req.URL.Query()
	q.Add("query", params.Get("query"))
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

func (s *RecipeService) GetRecipeInfo(id int, ctx context.Context) (map[string]any, bool, error) {
	cacheKey := strconv.Itoa(id)

	if cached, err := s.cache.Get(cacheKey); err == nil {
		var cachedResponse map[string]any
		if err := json.Unmarshal([]byte(cached), &cachedResponse); err == nil {
			return cachedResponse, true, nil
		}
	}

	url := fmt.Sprintf("/recipes/%d/information", id)
	req, err := s.newRequest("GET", url)
	if err != nil {
		return nil, false, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, false, err
	}
	defer res.Body.Close()

	var response map[string]any
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, false, err
	}

	responseJSON, err := json.Marshal(response)
	if err == nil {
		s.cache.Set(cacheKey, responseJSON, 120*time.Minute)
	}

	return response, false, nil
}

func (s *RecipeService) GetRecipeInfoBulk(ids []int, ctx context.Context) map[string]any {
	panic("unimplemented")
}
