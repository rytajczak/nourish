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
	Email        string   `json:"email"`
	Username     string   `json:"username"`
	Provider     string   `json:"provider"`
	Picture      string   `json:"picture"`
	Diet         string   `json:"diet"`
	Calories     int      `json:"calories"`
	Protein      int      `json:"protein"`
	Carbs        int      `json:"carbs"`
	Fat          int      `json:"fat"`
	Intolerances []string `json:"intolerances"`
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

func (s *ApiServer) handlePing(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, nil)
}

func (s *ApiServer) handleSignup(w http.ResponseWriter, r *http.Request) {
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

func (s *ApiServer) handleGetMe(w http.ResponseWriter, r *http.Request) {
	me, err := s.svc.GetMe(context.Background(), r.Header.Get("email"))
	if err != nil {
		WriteJSON(w, http.StatusNotFound, err.Error())
		return
	}

	WriteJSON(w, http.StatusOK, me)
}

// Start API Server
func (s *ApiServer) Start(listenAddr string) error {
	m := http.NewServeMux()

	m.HandleFunc("GET /v1/users/ping", s.handlePing)

	m.HandleFunc("POST /v1/users/signup", s.VerifyIDToken(s.handleSignup))
	m.HandleFunc("GET /v1/users/me", s.VerifyIDToken(s.handleGetMe))

	fmt.Println("Starting API Server on", listenAddr)
	return http.ListenAndServe(listenAddr, m)
}
