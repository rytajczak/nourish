package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateProfile(*Profile) error
	GetProfileByID(string) (*Profile, error)
	UpdateProfile(*Profile) error
	DeleteProfile(string) error
}

type PostgresStore struct {
	db *sql.DB
}

func (s *PostgresStore) CreateProfile(profile *Profile) error {
	return nil
}

func (s *PostgresStore) GetProfileByID(id string) (*Profile, error) {
	return nil, nil
}

func (s *PostgresStore) UpdateProfile(profile *Profile) error {
	return nil
}

func (s *PostgresStore) DeleteProfile(id string) error {
	return nil
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=user dbname=postgres host=user-db password=password sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return &PostgresStore{
		db: db,
	}, nil
}
