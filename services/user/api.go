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
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		WriteJSON(w, http.StatusOK, map[string]string{"msg": "ok"})
	})

	http.HandleFunc("/user/signup", func(w http.ResponseWriter, r *http.Request) {
		var d SignUpDetails
		if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		profile, err := s.store.CreateUser(&d)
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		WriteJSON(w, http.StatusOK, profile)
	})

	http.HandleFunc("/profile/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		switch r.Method {
		case "GET":
			profile, err := s.store.GetProfileByID(id)
			if err != nil {
				WriteJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
				return
			}
			WriteJSON(w, http.StatusOK, profile)
			break
		case "DELETE":
			s.store.DeleteUser(id)
		}
	})

	log.Printf("listening on %s", listenAddr)
	return http.ListenAndServe(listenAddr, nil)
}
