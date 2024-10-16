package main

import (
	"context"
	"encoding/json"
	"log"
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

func (s *ApiServer) Start(listenAddr string) error {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{"msg": "ok"})
	})

	http.HandleFunc("/recipes", func(w http.ResponseWriter, r *http.Request) {
		data := s.svc.GetRandomRecipes(context.Background())
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	})

	log.Printf("listening on %s", listenAddr)
	return http.ListenAndServe(listenAddr, nil)
}
