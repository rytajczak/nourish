package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

type ApiServer struct {
	svc Service
}

func NewApiServer(svc Service) *ApiServer {
	return &ApiServer{
		svc: svc,
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func (s *ApiServer) handlePing(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, 200, nil)
}

func (s *ApiServer) handleSearch(w http.ResponseWriter, r *http.Request) {
	recipes := s.svc.SearchRecipes(r.URL.Query(), context.Background())
	WriteJSON(w, 200, recipes)
}

func (s *ApiServer) handleGetInfo(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]any{"error": "'id' must be a number"})
	}

	result, _, err := s.svc.GetRecipeInfo(id, context.Background())

	WriteJSON(w, http.StatusOK, result)
}

func (s *ApiServer) Start(listenAddr string) error {
	m := http.NewServeMux()

	m.HandleFunc("GET /v1/recipes/ping", s.handlePing)

	m.HandleFunc("GET /v1/recipes/search", s.handleSearch)
	m.HandleFunc("GET /v1/recipes/{id}/info", s.handleGetInfo)
	m.HandleFunc("GET /v1/recipes/info-bulk", s.handleSearch)

	return http.ListenAndServe(listenAddr, m)
}
