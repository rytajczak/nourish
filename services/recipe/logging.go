package main

import (
	"context"
	"net/url"
	"time"

	"github.com/charmbracelet/log"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &LoggingService{
		next: next,
	}
}

func (s *LoggingService) SearchRecipes(params url.Values, ctx context.Context) (map[string]any, error) {
	defer func(start time.Time) {
		log.Info("Search recipes finished", "took", time.Since(start))
	}(time.Now())
	return s.next.SearchRecipes(params, ctx)
}

func (s *LoggingService) GetRecipeInfo(id int, ctx context.Context) (map[string]any, error) {
	defer func(start time.Time) {
		log.Info("Get recipe finished", "took", time.Since(start))
	}(time.Now())
	return s.next.GetRecipeInfo(id, ctx)
}

func (s *LoggingService) GetRecipeInfoBulk(ids string, ctx context.Context) ([]map[string]any, error) {
	defer func(start time.Time) {
		log.Info("Get bulk recipes finished", "took", time.Since(start))
	}(time.Now())
	return s.next.GetRecipeInfoBulk(ids, ctx)
}
