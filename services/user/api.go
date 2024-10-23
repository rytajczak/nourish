package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ApiServer struct {
	store Storage
}

func NewApiServer(store Storage) *ApiServer {
	return &ApiServer{
		store: store,
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func (s *ApiServer) Start(listenAddr string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		WriteJSON(w, http.StatusOK, map[string]any{"msg": "ok"})
	})

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		WriteJSON(w, http.StatusOK, map[string]any{"msg": "ok"})
	})

	log.Printf("listening on port %s", listenAddr)
	return http.ListenAndServe(listenAddr, mux)
}
