package main

import (
	"context"
	"net/url"

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

func (s *LoggingService) SearchRecipes(params url.Values, ctx context.Context) map[string]any {
	panic("unimplemented")
}

func (s *LoggingService) GetRecipeInfo(id int, ctx context.Context) (map[string]any, bool, error) {
	log.Info("getting recipe", "id", id)
	result, cached, err := s.next.GetRecipeInfo(id, ctx)
	if err != nil {
		log.Error("Failed to fetch recipe", "error", err)
		return result, cached, nil
	}

	if cached {
		log.Info("retrieved recipe from cache", "id", id)
	} else {
		log.Info("retrieved recipe from API", "id", id)
	}
	return result, cached, nil
}

func (s *LoggingService) GetRecipeInfoBulk(ids []int, ctx context.Context) map[string]any {
	panic("unimplemented")
}
