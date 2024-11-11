package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"user/repository"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ApiServer struct {
	queries *repository.Queries
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Provider string `json:"provider"`
	Picture  string `json:"picture"`
}

type SpoonUserConnectResponse struct {
	Status              string `json:"status"`
	Username            string `json:"username"`
	SpoonacularPassword string `json:"spoonacularPassword"`
	Hash                string `json:"hash"`
}

type UserProfileResponse struct {
	User         repository.User            `json:"user"`
	DailyGoal    repository.GetDailyGoalRow `json:"dailyGoal"`
	Intolerances []string                   `json:"intolerances"`
	Dislikes     []string                   `json:"dislikes"`
	Likes        []repository.LikedRecipe   `json:"likes"`
}

type UpdateUserDailyGoalRequest struct {
	Calories int `json:"calories"`
	Carbs    int `json:"carbs"`
	Protein  int `json:"protein"`
	Fat      int `json:"fat"`
}

type UpdateUserIntolerancesRequest struct {
	Intolerances []string `json:"intolerances"`
}

func NewApiServer(pool *pgxpool.Pool) *ApiServer {
	return &ApiServer{
		queries: repository.New(pool),
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func (s *ApiServer) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("https://%s/users/connect", os.Getenv("API_HOST"))
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	req.Header.Add("x-rapidapi-key", os.Getenv("API_KEY"))
	req.Header.Add("x-rapidapi-host", os.Getenv("API_HOST"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer resp.Body.Close()

	spoonResp, err := io.ReadAll(resp.Body)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	var spoonBody SpoonUserConnectResponse
	if err := json.Unmarshal(spoonResp, &spoonBody); err != nil {
		WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if spoonBody.Status != "success" {
		WriteJSON(w, http.StatusBadRequest, spoonBody.Status)
		return
	}

	var body CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := s.queries.CreateUser(context.Background(), repository.CreateUserParams{
		Username: body.Username,
		Email:    body.Email,
		Provider: body.Provider,
		Picture:  pgtype.Text{String: body.Picture, Valid: true},
	})
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = s.queries.CreateSpoonCredential(context.Background(), repository.CreateSpoonCredentialParams{
		UserID:   user.ID,
		Username: spoonBody.Username,
		Password: spoonBody.SpoonacularPassword,
		Hash:     spoonBody.Hash,
	})
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	WriteJSON(w, http.StatusOK, user)
}

func (s *ApiServer) handleGetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()

	var profile UserProfileResponse
	var userErr, dailyGoalErr, intolerancesErr, dislikesErr, likesErr error

	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		defer wg.Done()
		profile.User, userErr = s.queries.GetUserById(ctx, pgtype.UUID{Bytes: id, Valid: true})
	}()

	go func() {
		defer wg.Done()
		profile.DailyGoal, dailyGoalErr = s.queries.GetDailyGoal(ctx, pgtype.UUID{Bytes: id, Valid: true})
	}()

	go func() {
		defer wg.Done()
		profile.Intolerances, intolerancesErr = s.queries.GetIntolerances(ctx, pgtype.UUID{Bytes: id, Valid: true})
	}()

	go func() {
		defer wg.Done()
		profile.Dislikes, dislikesErr = s.queries.GetDislikedIngredients(ctx, pgtype.UUID{Bytes: id, Valid: true})
	}()

	go func() {
		defer wg.Done()
		profile.Likes, likesErr = s.queries.GetLikedRecipes(ctx, pgtype.UUID{Bytes: id, Valid: true})
	}()

	wg.Wait()

	if userErr != nil || dailyGoalErr != nil || intolerancesErr != nil || dislikesErr != nil || likesErr != nil {
		fmt.Println(userErr, dailyGoalErr, intolerancesErr, dislikesErr, likesErr)
		WriteJSON(w, http.StatusInternalServerError, "Failed to fetch user profile")
		return
	}

	WriteJSON(w, http.StatusOK, profile)
}

func (s *ApiServer) handleGetIdByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	user, err := s.queries.GetUserByEmail(context.Background(), email)
	if err != nil {
		WriteJSON(w, http.StatusNotFound, map[string]string{"msg": "no user found"})
		return
	}

	WriteJSON(w, http.StatusOK, user.ID)
}

// Update User Preferences
func (s *ApiServer) handleUpdateUserDailyGoal(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	var body UpdateUserDailyGoalRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	s.queries.UpdateDailyGoal(context.Background(), repository.UpdateDailyGoalParams{
		UserID:   pgtype.UUID{Bytes: id, Valid: true},
		Calories: pgtype.Int4{Int32: int32(body.Calories), Valid: true},
		Carbs:    pgtype.Int4{Int32: int32(body.Carbs), Valid: true},
		Protein:  pgtype.Int4{Int32: int32(body.Protein), Valid: true},
		Fat:      pgtype.Int4{Int32: int32(body.Fat), Valid: true},
	})

	WriteJSON(w, http.StatusOK, nil)
}

func (s *ApiServer) Start(listenAddr string) error {
	m := http.NewServeMux()

	// User Account Management
	m.HandleFunc("POST /users/", s.handleCreateUser)
	m.HandleFunc("GET /users/{id}", s.handleGetUser)
	m.HandleFunc("GET /users/getIdByEmail", s.handleGetIdByEmail)

	// User Preferences
	m.HandleFunc("PUT /users/{id}/dailyGoal", s.handleUpdateUserDailyGoal)
	// m.HandleFunc("PUT /users/{id}/intolerances", s.handleUpdateUserIntolerances)

	// // User Recipes
	// m.HandleFunc("GET /users/{id}/recipes", s.handleGetUserRecipes)
	// m.HandleFunc("PATCH /users/{id}/recipes", s.handleUpdateUserRecipes)

	return http.ListenAndServe(listenAddr, m)
}
