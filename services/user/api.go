package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"user/repository"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/api/idtoken"
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

func (s *ApiServer) VerifyIDToken(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			WriteJSON(w, http.StatusUnauthorized, "Authorization header is required")
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			WriteJSON(w, http.StatusUnauthorized, "Bearer token is required")
			return
		}

		ctx := context.Background()
		payload, err := idtoken.Validate(ctx, token, os.Getenv("GOOGLE_CLIENT_ID"))
		if err != nil {
			WriteJSON(w, http.StatusUnauthorized, err.Error())
			return
		}

		fmt.Println(payload)

		next.ServeHTTP(w, r)
	})
}

func (s *ApiServer) handleHealth(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, map[string]string{"msg": "ok"})
}

// Create User
func (s *ApiServer) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Starting handleCreateUser")
	url := fmt.Sprintf("https://%s/users/connect", os.Getenv("API_HOST"))
	fmt.Printf("Making request to URL: %s\n", url)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Printf("Failed to create request: %v\n", err)
		WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	req.Header.Add("x-rapidapi-key", os.Getenv("API_KEY"))
	req.Header.Add("x-rapidapi-host", os.Getenv("API_HOST"))
	fmt.Printf("Added headers - Host: %s\n", os.Getenv("API_HOST"))

	client := &http.Client{}
	fmt.Println("Making HTTP request to Spoonacular")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("HTTP request failed: %v\n", err)
		WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer resp.Body.Close()

	spoonResp, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %v\n", err)
		WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Printf("Received response from Spoonacular: %s\n", string(spoonResp))

	var spoonBody SpoonUserConnectResponse
	if err := json.Unmarshal(spoonResp, &spoonBody); err != nil {
		fmt.Printf("Failed to unmarshal Spoonacular response: %v\n", err)
		WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if spoonBody.Status != "success" {
		fmt.Printf("Spoonacular returned non-success status: %s\n", spoonBody.Status)
		WriteJSON(w, http.StatusBadRequest, spoonBody.Status)
		return
	}
	fmt.Println("Successfully connected to Spoonacular")

	var body CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Printf("Failed to decode request body: %v\n", err)
		WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Printf("Received user creation request for email: %s\n", body.Email)

	user, err := s.queries.CreateUser(context.Background(), repository.CreateUserParams{
		Username: body.Username,
		Email:    body.Email,
		Provider: body.Provider,
		Picture:  pgtype.Text{String: body.Picture, Valid: true},
	})
	if err != nil {
		fmt.Printf("Failed to create user in database: %v\n", err)
		WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Printf("Created user with ID: %v\n", user.ID)

	_, err = s.queries.CreateSpoonCredential(context.Background(), repository.CreateSpoonCredentialParams{
		UserID:   user.ID,
		Username: spoonBody.Username,
		Password: spoonBody.SpoonacularPassword,
		Hash:     spoonBody.Hash,
	})
	if err != nil {
		fmt.Printf("Failed to create Spoonacular credentials: %v\n", err)
		WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("Successfully created Spoonacular credentials")

	WriteJSON(w, http.StatusOK, user)
	fmt.Println("User creation completed successfully")
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

	m.HandleFunc("GET /users/health", s.handleHealth)

	// User Account Management
	m.HandleFunc("POST /users/", s.VerifyIDToken(s.handleCreateUser))
	m.HandleFunc("GET /users/me", s.VerifyIDToken(s.handleGetUser))

	// User Preferences
	m.HandleFunc("PUT /users/me/dailyGoal", s.VerifyIDToken(s.handleUpdateUserDailyGoal))

	return http.ListenAndServe(listenAddr, m)
}
