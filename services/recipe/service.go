package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/charmbracelet/log"
)

type Service interface {
	SearchRecipes(params url.Values, ctx context.Context) (map[string]any, error)
	GetRecipeInfo(id int, ctx context.Context) (map[string]any, error)
	GetRecipeInfoBulk(ids string, ctx context.Context) ([]map[string]any, error)
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

func (s *RecipeService) SearchRecipes(params url.Values, ctx context.Context) (map[string]any, error) {
	req, err := s.newRequest("GET", "/recipes/complexSearch")
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for key, values := range params {
		for _, value := range values {
			q.Add(key, value)
		}
	}
	q.Add("addRecipeNutrition", "true")
	q.Add("instructionsRequired", "true")
	q.Add("number", "30")
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response map[string]any
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *RecipeService) GetRecipeInfo(id int, ctx context.Context) (map[string]any, error) {
	cacheKey := fmt.Sprintf("recipe:%d", id)

	if cached, err := s.cache.Get(cacheKey); err == nil {
		var cachedResponse map[string]any
		if err := json.Unmarshal([]byte(cached), &cachedResponse); err == nil {
			return cachedResponse, nil
		}
	}

	url := fmt.Sprintf("/recipes/%d/information", id)
	req, err := s.newRequest("GET", url)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("includeNutrition", "true")
	req.URL.RawQuery = q.Encode()

	log.Info(req.URL)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response map[string]any
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	responseJSON, err := json.Marshal(response)
	if err == nil {
		s.cache.Set(cacheKey, responseJSON, 120*time.Minute)
	}

	return response, nil
}

func (s *RecipeService) GetRecipeInfoBulk(idsCSV string, ctx context.Context) ([]map[string]any, error) {
	var results []map[string]any
	var uncachedIDs []string

	idStrings := strings.Split(idsCSV, ",")
	for _, idStr := range idStrings {
		id := strings.TrimSpace(idStr)
		cacheKey := fmt.Sprintf("recipe:%s", id)
		cachedData, err := s.cache.Get(cacheKey)
		if err == nil && cachedData != "" {
			var result map[string]any
			if err := json.Unmarshal([]byte(cachedData), &result); err == nil {
				results = append(results, result)
				continue
			}
		}
		uncachedIDs = append(uncachedIDs, id)
	}

	if len(uncachedIDs) > 0 {
		idCSV := strings.Join(uncachedIDs, ",")

		url := fmt.Sprintf("/recipes/informationBulk?ids=%s", idCSV)
		req, err := s.newRequest("GET", url)
		if err != nil {
			return nil, err
		}

		q := req.URL.Query()
		q.Add("includeNutrition", "true")
		req.URL.RawQuery = q.Encode()

		log.Info(req.URL)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		var apiResults []map[string]any
		if err := json.NewDecoder(res.Body).Decode(&apiResults); err != nil {
			return nil, err
		}

		for _, result := range apiResults {
			id := int(result["id"].(float64))
			results = append(results, result)
			data, _ := json.Marshal(result)
			cacheKey := fmt.Sprintf("recipe:%d", id)
			s.cache.Set(cacheKey, string(data), 0)
		}
	}

	return results, nil
}
