package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"user/repository"

	"github.com/jackc/pgx/v5/pgtype"
)

type ApiServer struct {
	repo *repository.Queries
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

func NewApiServer(repo *repository.Queries) *ApiServer {
	return &ApiServer{repo: repo}
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

	user, err := s.repo.CreateUser(context.Background(), repository.CreateUserParams{
		Username: body.Username,
		Email:    body.Email,
		Provider: body.Provider,
		Picture:  pgtype.Text{String: body.Picture, Valid: true},
	})
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = s.repo.CreateSpoonCredential(context.Background(), repository.CreateSpoonCredentialParams{
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
	id := r.PathValue("id")
	WriteJSON(w, http.StatusOK, id)
}

func (s *ApiServer) Start(listenAddr string) error {
	m := http.NewServeMux()

	m.HandleFunc("POST /users/", s.handleCreateUser)
	m.HandleFunc("GET /users/{id}", s.handleGetUser)

	return http.ListenAndServe(listenAddr, m)
}
