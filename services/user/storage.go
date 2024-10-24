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

	for _, intolerance := range details.Intolerances {
		_, err := tx.Exec(`
			INSERT INTO profile_intolerance (profile_id, intolerance_id)
			SELECT $1, i.id
			FROM intolerance i
			WHERE i.name = $2
		`, id, intolerance)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("failed to insert intolerance: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	log.Printf("created user '%s'", details.Username)
	return s.GetProfileByID(id)
}

func (s *PostgresStore) GetProfileByID(id string) (*Profile, error) {
	var profile Profile

	query := `
		SELECT 
				p.id, p.username, p.first_name, p.last_name, p.diet,
				COALESCE(ARRAY_AGG(i.name::text), '{}') AS intolerances
		FROM profile p
		LEFT JOIN profile_intolerance pi ON p.id = pi.profile_id
		LEFT JOIN intolerance i ON pi.intolerance_id = i.id
		WHERE p.id = $1
		GROUP BY p.id;
    `
	err := s.db.Get(&profile, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get profile: %w", err)
	}

	return &profile, nil
}

func (s *PostgresStore) UpdateProfile(profile *Profile) error {
	return nil
}

func (s *PostgresStore) DeleteUser(id string) error {
	tx := s.db.MustBegin()

	_, err := tx.Exec("DELETE FROM auth WHERE id = $1", id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete user: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

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
