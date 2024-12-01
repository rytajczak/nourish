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
	CreateUser(request CreateUserRequest, ctx context.Context) (map[string]any, error)
	GetMe(email string, ctx context.Context) (*UserResponse, error)
	UpdateProfile(email string, profile map[string]any, ctx context.Context) (map[string]any, error)
	UpdateIntolerances(email string, intolerances []string, ctx context.Context) ([]string, error)
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
func (s *UserService) CreateUser(request CreateUserRequest, ctx context.Context) (map[string]any, error) {
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
		Username: request.Username,
		Email:    request.Email,
		Provider: request.Provider,
		Picture:  pgtype.Text{String: request.Picture, Valid: true},
		Diet:     pgtype.Text{String: request.Profile.Diet, Valid: true},
		Calories: pgtype.Int4{Int32: int32(request.Profile.Calories), Valid: true},
		Protein:  pgtype.Int4{Int32: int32(request.Profile.Protein), Valid: true},
		Carbs:    pgtype.Int4{Int32: int32(request.Profile.Carbs), Valid: true},
		Fat:      pgtype.Int4{Int32: int32(request.Profile.Carbs), Valid: true},
	})

	profile, err := s.queries.GetUserProfile(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	intolerances := []string{}
	if len(request.Intolerances) > 0 {
		intolerances, err = s.UpdateIntolerances(user.Email, request.Intolerances, ctx)
		if err != nil {
			return nil, err
		}

		if intolerances == nil {
			intolerances = []string{}
		}
	}

	_, err = s.queries.CreateSpoonCredential(ctx, repository.CreateSpoonCredentialParams{
		UserID:   user.ID,
		Username: response.Username,
		Password: response.SpoonacularPassword,
		Hash:     response.Hash,
	})

	spoonCredential, err := s.queries.GetUsernameAndHash(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	result := map[string]any{
		"profile":         profile,
		"intolerances":    intolerances,
		"spoonCredential": spoonCredential,
	}

	return result, nil
}

// GetMe gets a user by their email
func (s *UserService) GetMe(email string, ctx context.Context) (*UserResponse, error) {
	user, err := s.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	profile, err := s.queries.GetUserProfile(ctx, email)
	if err != nil {
		return nil, err
	}

	spoonCredential, err := s.queries.GetUsernameAndHash(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	intolerances, _ := s.queries.GetUserIntolerances(ctx, email)
	if intolerances == nil {
		intolerances = []string{}
	}

	return &UserResponse{
		Profile:         profile,
		Intolerances:    intolerances,
		SavedRecipes:    []int{},
		SpoonCredential: spoonCredential,
	}, nil
}

func (s *UserService) UpdateProfile(email string, profile map[string]any, ctx context.Context) (map[string]any, error) {
	return nil, nil
}

func (s *UserService) UpdateIntolerances(email string, intolerances []string, ctx context.Context) ([]string, error) {
	user, err := s.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	s.queries.DeleteUserIntolerance(ctx, user.ID)

	for _, intolerance := range intolerances {
		_, err = s.queries.CreateUserIntolerance(ctx, repository.CreateUserIntoleranceParams{
			UserID: user.ID,
			Name:   intolerance,
		})
	}

	return intolerances, nil
}
