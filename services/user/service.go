package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"user/repository"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service interface {
	CreateUser(ctx context.Context, info CreateUserRequest) error
}

type UserService struct {
	queries *repository.Queries
	url     string
	host    string
	key     string
}

func NewUserService(host string, key string) Service {
	config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	config.MaxConns = 10
	config.MinConns = 1

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	url := fmt.Sprintf("https://%s", host)

	return &UserService{queries: repository.New(pool), url: url, host: host, key: key}
}

func (s *UserService) CreateUser(ctx context.Context, info CreateUserRequest) error {
	req, _ := http.NewRequest("POST", s.url+"/users/connect", nil)
	req.Header.Add("x-rapidapi-key", s.key)
	req.Header.Add("x-rapidapi-host", s.host)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var response SpoonUserConnectResponse
	json.Unmarshal(body, &response)

	user, err := s.queries.CreateUser(ctx, repository.CreateUserParams{
		Username: info.Username,
		Email:    info.Email,
		Provider: info.Provider,
		Picture:  pgtype.Text{String: info.Picture, Valid: true},
	})

	_, err = s.queries.CreateSpoonCredential(ctx, repository.CreateSpoonCredentialParams{
		UserID:   user.ID,
		Username: response.Username,
		Password: response.SpoonacularPassword,
		Hash:     response.Hash,
	})

	return nil
}

func (s *UserService) GetUser(ctx context.Context, id string) (*repository.User, error) {
	return nil, nil
}

func (s *UserService) UpdateUserPreferences(ctx context.Context, user *repository.User) error {
	return nil
}
