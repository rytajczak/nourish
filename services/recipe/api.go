package main

import (
	"context"
	"encoding/json"
	"net/http"
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
	params := r.URL.Query()
	query := params.Get("query")
	recipes := s.svc.SearchRecipes(query, context.Background())
	WriteJSON(w, 200, recipes)
}

func (s *ApiServer) Start(listenAddr string) error {
	m := http.NewServeMux()

	m.HandleFunc("GET /v1/recipes/ping", s.handlePing)
	m.HandleFunc("GET /v1/recipes/search", s.handleSearch)

	return http.ListenAndServe(listenAddr, m)
}
