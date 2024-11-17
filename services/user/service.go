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
	CreateUser(ctx context.Context, info CreateUserRequest) (*repository.User, error)
	GetUser(ctx context.Context, email string) (*repository.User, error)
	UpdateUserPreferences(ctx context.Context, email string, info *UpdateUserPreferencesRequest) (*repository.User, error)
	GetUserIntolerances(ctx context.Context, email string) ([]string, error)
	UpdateUserIntolerances(ctx context.Context, email string, intolerances []string) ([]string, error)
}

type UserService struct {
	queries *repository.Queries
	url     string
	host    string
	key     string
}

type SpoonUserConnectResponse struct {
	Status              string `json:"status"`
	Username            string `json:"username"`
	SpoonacularPassword string `json:"spoonacularPassword"`
	Hash                string `json:"hash"`
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

// CreateUser creates a new user and connects them to their spoon account
func (s *UserService) CreateUser(ctx context.Context, info CreateUserRequest) (*repository.User, error) {
	req, _ := http.NewRequest("POST", s.url+"/users/connect", nil)
	req.Header.Add("x-rapidapi-key", s.key)
	req.Header.Add("x-rapidapi-host", s.host)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
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

	return &user, nil
}

// GetUser gets a user by their email
func (s *UserService) GetUser(ctx context.Context, email string) (*repository.User, error) {
	user, err := s.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUserPreferences updates a user's preferences
func (s *UserService) UpdateUserPreferences(ctx context.Context, email string, info *UpdateUserPreferencesRequest) (*repository.User, error) {
	updatedUser, err := s.queries.UpdateUserPreferences(ctx, repository.UpdateUserPreferencesParams{
		Email:    email,
		Diet:     pgtype.Text{String: info.Diet, Valid: true},
		Calories: pgtype.Int4{Int32: int32(info.Calories), Valid: true},
		Protein:  pgtype.Int4{Int32: int32(info.Protein), Valid: true},
		Carbs:    pgtype.Int4{Int32: int32(info.Carbs), Valid: true},
		Fat:      pgtype.Int4{Int32: int32(info.Fat), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (s *UserService) GetUserIntolerances(ctx context.Context, email string) ([]string, error) {
	intolerances, err := s.queries.GetUserIntolerances(ctx, email)
	if err != nil {
		return nil, err
	}
	return intolerances, nil
}

func (s *UserService) UpdateUserIntolerances(ctx context.Context, email string, intolerances []string) ([]string, error) {
	fmt.Println("updating intolerances for user", email)
	user, err := s.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	fmt.Println("deleting intolerances for user", email)
	_, err = s.queries.DeleteUserIntolerance(ctx, user.ID)
	if err != nil {
		fmt.Println("error deleting intolerances", err)
	}

	fmt.Println("adding intolerances for user", email)
	for _, intolerance := range intolerances {
		fmt.Println("adding intolerance", intolerance, "for user", user.ID)
		_, err = s.queries.CreateUserIntolerance(ctx, repository.CreateUserIntoleranceParams{
			UserID: user.ID,
			Name:   intolerance,
		})
	}
	fmt.Println("done adding intolerances")

	return intolerances, nil
}
