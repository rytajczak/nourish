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

func (s *LoggingService) CreateUser(request CreateUserRequest, ctx context.Context) (map[string]any, error) {
	s.logger.Info("Creating user", "request", request)

	result, err := s.next.CreateUser(request, ctx)
	if err != nil {
		s.logger.Error("Failed to create user", "err", err, "request", request)
		return nil, err
	}
	s.logger.Info("Create user finished", "result", result)

	return result, nil
}

func (s *LoggingService) GetMe(email string, ctx context.Context) (*UserResponse, error) {
	s.logger.Info("Getting user", "email", email)
	return s.next.GetMe(email, ctx)
}

func (s *LoggingService) UpdateIntolerances(email string, intolerances []string, ctx context.Context) ([]string, error) {
	panic("unimplemented")
}

func (s *LoggingService) UpdateProfile(email string, profile map[string]any, ctx context.Context) (map[string]any, error) {
	panic("unimplemented")
}
