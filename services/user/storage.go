package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateUser(*SignUpDetails) (*Profile, error)
	GetProfileByID(string) (*Profile, error)
	UpdateProfile(*Profile) error
	DeleteUser(string) error
}

type PostgresStore struct {
	db *sqlx.DB
}

func (s *PostgresStore) CreateUser(details *SignUpDetails) (*Profile, error) {
	log.Print(details)
	id := uuid.New().String()

	tx := s.db.MustBegin()
	if _, err := tx.Exec("INSERT INTO auth (id, email, provider, password) VALUES ($1, $2, $3, $4)", id, details.Email, details.Provider, details.Password); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to insert into auth: %s", err.Error())
	}

	query := `INSERT INTO profile (id, username, first_name, last_name, diet)
						VALUES ($1, $2, $3, $4, $5)`

	_, err := tx.Exec(query, id, details.Username, details.FirstName, details.LastName, details.Diet)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	log.Printf("created user '%s'", details.Username)

	var createdProfile Profile
	if err := s.db.Get(&createdProfile, "SELECT * FROM profile WHERE id=$1", id); err != nil {
		return nil, err
	}

	return &createdProfile, nil
}

func (s *PostgresStore) GetProfileByID(id string) (*Profile, error) {
	var createdProfile Profile
	if err := s.db.Get(&createdProfile, "SELECT * FROM profile WHERE id=$1", id); err != nil {
		return nil, err
	}

	return &createdProfile, nil
}

func (s *PostgresStore) UpdateProfile(profile *Profile) error {
	return nil
}

func (s *PostgresStore) DeleteUser(id string) error {
	return nil
}

func NewPostgresStore() (*PostgresStore, error) {
	db, err := sqlx.Connect("postgres", "user=user dbname=user password=password sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	return &PostgresStore{
		db: db,
	}, nil
}
