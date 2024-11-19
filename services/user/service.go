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
	CreateUser(ctx context.Context, info CreateUserRequest) (map[string]any, error)
	GetMe(ctx context.Context, email string) (map[string]any, error)
	UpdateProfile(ctx context.Context, email string, info *UpdateUserPreferencesRequest) (*repository.UpdateUserProfileRow, error)
	UpdateIntolerances(ctx context.Context, email string, intolerances []string) ([]string, error)
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
func (s *UserService) CreateUser(ctx context.Context, info CreateUserRequest) (map[string]any, error) {
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
		Diet:     pgtype.Text{String: info.Diet, Valid: true},
		Calories: pgtype.Int4{Int32: int32(info.Calories), Valid: true},
		Protein:  pgtype.Int4{Int32: int32(info.Protein), Valid: true},
		Carbs:    pgtype.Int4{Int32: int32(info.Carbs), Valid: true},
		Fat:      pgtype.Int4{Int32: int32(info.Fat), Valid: true},
	})

	profile, err := s.queries.GetUserProfile(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	var intolerances []string
	if len(info.Intolerances) > 0 {
		intolerances, err = s.UpdateIntolerances(ctx, user.Email, info.Intolerances)
		if err != nil {
			return nil, err
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

	return map[string]any{"profile": profile, "intolerances": intolerances, "spoonCredential": spoonCredential}, nil
}

// GetMe gets a user by their email
func (s *UserService) GetMe(ctx context.Context, email string) (map[string]any, error) {
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

	intolerances, err := s.queries.GetUserIntolerances(ctx, email)
	if err != nil {
		return nil, err
	}

	if intolerances == nil {
		intolerances = []string{}
	}

	return map[string]any{"profile": profile, "intolerances": intolerances, "spoonCredential": spoonCredential}, nil
}

// UpdateUserPreferences updates a user's preferences
func (s *UserService) UpdateProfile(ctx context.Context, email string, info *UpdateUserPreferencesRequest) (*repository.UpdateUserProfileRow, error) {
	updatedProfile, err := s.queries.UpdateUserProfile(ctx, repository.UpdateUserProfileParams{
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

	return &updatedProfile, nil
}

func (s *UserService) UpdateIntolerances(ctx context.Context, email string, intolerances []string) ([]string, error) {
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
