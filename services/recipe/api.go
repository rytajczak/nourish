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
	WriteJSON(w, http.StatusOK, nil)
}

func (s *ApiServer) handleSearch(w http.ResponseWriter, r *http.Request) {
	recipes, err := s.svc.SearchRecipes(r.URL.Query(), context.Background())
	if err != nil {
		WriteJSON(w, http.StatusBadGateway, nil)
		return
	}

	WriteJSON(w, http.StatusOK, recipes)
}

func (s *ApiServer) handleGetInfo(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]any{"error": "'id' must be a number"})
		return
	}

	result, err := s.svc.GetRecipeInfo(id, context.Background())

	WriteJSON(w, http.StatusOK, result)
}

func (s *ApiServer) handleGetBulkInfo(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	ids := queryParams.Get("ids")

	result, err := s.svc.GetRecipeInfoBulk(ids, context.Background())
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, nil)
	}

	WriteJSON(w, http.StatusOK, result)
}

func (s *ApiServer) Start(listenAddr string) error {
	m := http.NewServeMux()

	m.HandleFunc("GET /v1/recipes/ping", s.handlePing)

	m.HandleFunc("GET /v1/recipes/search", s.handleSearch)
	m.HandleFunc("GET /v1/recipes/{id}/info", s.handleGetInfo)
	m.HandleFunc("GET /v1/recipes/info-bulk", s.handleGetBulkInfo)

	return http.ListenAndServe(listenAddr, m)
}
