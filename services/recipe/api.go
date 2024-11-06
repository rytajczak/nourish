package main

import (
	"context"
	"encoding/json"
	"log"
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

func (s *ApiServer) Start(listenAddr string) error {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		WriteJSON(w, http.StatusOK, map[string]string{"msg": "ok"})
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		query := params.Get("query")
		log.Printf("processing search for '%s'", query)
		recipes := s.svc.SearchRecipes(query, context.Background())
		WriteJSON(w, 200, recipes)
	})

	http.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		strId := r.PathValue("id")
		id, err := strconv.Atoi(strId)
		if err != nil {
			log.Println("Couldn't parse id")
		}

		recipeInfo := s.svc.GetRecipeById(id, context.Background())
		WriteJSON(w, 200, recipeInfo)
	})

	log.Printf("listening on %s", listenAddr)
	return http.ListenAndServe(listenAddr, nil)
}
