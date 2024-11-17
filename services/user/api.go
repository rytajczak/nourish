package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"google.golang.org/api/idtoken"
)

type ApiServer struct {
	svc Service
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Provider string `json:"provider"`
	Picture  string `json:"picture"`
}

type UpdateUserPreferencesRequest struct {
	Calories     int      `json:"calories"`
	Protein      int      `json:"protein"`
	Carbs        int      `json:"carbs"`
	Fat          int      `json:"fat"`
	Diet         string   `json:"diet"`
	Intolerances []string `json:"intolerances"`
}

func NewApiServer(svc Service) *ApiServer {
	return &ApiServer{svc: svc}
}

// Write JSON Response
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// Verify ID Token
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

		email := payload.Claims["email"].(string)
		r.Header.Set("email", email)

		next.ServeHTTP(w, r)
	})
}

func (s *ApiServer) handleHealth(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, map[string]string{"msg": "ok"})
}

func (s *ApiServer) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var body CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := s.svc.CreateUser(context.Background(), body)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	WriteJSON(w, http.StatusOK, user)
}

func (s *ApiServer) handleGetUser(w http.ResponseWriter, r *http.Request) {
	user, err := s.svc.GetUser(context.Background(), r.Header.Get("email"))
	if err != nil {
		WriteJSON(w, http.StatusNotFound, map[string]string{"error": "user not found"})
		return
	}

	WriteJSON(w, http.StatusOK, user)
}

func (s *ApiServer) handleUpdateUserPreferences(w http.ResponseWriter, r *http.Request) {
	var body UpdateUserPreferencesRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	s.svc.UpdateUserPreferences(context.Background(), r.Header.Get("email"), &body)
	s.svc.UpdateUserIntolerances(context.Background(), r.Header.Get("email"), body.Intolerances)

	WriteJSON(w, http.StatusOK, body)
}

// Start API Server
func (s *ApiServer) Start(listenAddr string) error {
	m := http.NewServeMux()

	m.HandleFunc("GET /v1/users/health", s.handleHealth)

	m.HandleFunc("POST /v1/users/", s.VerifyIDToken(s.handleCreateUser))
	m.HandleFunc("GET /v1/users/me", s.VerifyIDToken(s.handleGetUser))
	m.HandleFunc("PUT /v1/users/me/preferences", s.VerifyIDToken(s.handleUpdateUserPreferences))

	fmt.Println("Starting API Server on", listenAddr)
	return http.ListenAndServe(listenAddr, m)
}
