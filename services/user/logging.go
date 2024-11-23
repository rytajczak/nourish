package main

import (
	"context"

	"github.com/charmbracelet/log"
)

type LoggingService struct {
	next   Service
	logger *log.Logger
}

func NewLoggingService(next Service, logger *log.Logger) Service {
	return &LoggingService{
		next:   next,
		logger: logger,
	}
}

func (s *LoggingService) CreateUser(ctx context.Context, request CreateUserRequest) (map[string]any, error) {
	s.logger.Info("Creating user", "request", request)

	result, err := s.next.CreateUser(ctx, request)
	if err != nil {
		s.logger.Error("Failed to create user", "err", err, "request", request)
		return nil, err
	}
	s.logger.Info("Create user finished", "result", result)

	return result, nil
}

func (s *LoggingService) GetMe(ctx context.Context, email string) (map[string]any, error) {
	return s.next.GetMe(ctx, email)
}

func (s *LoggingService) UpdateIntolerances(ctx context.Context, email string, intolerances []string) ([]string, error) {
	panic("unimplemented")
}

func (s *LoggingService) UpdateProfile(ctx context.Context, email string, profile map[string]any) (map[string]any, error) {
	panic("unimplemented")
}
