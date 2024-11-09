package main

import (
	"context"
	"log"
	"os"
	"user/repository"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close(ctx)

	repo := repository.New(conn)

	repo.CreateUser(ctx, repository.CreateUserParams{
		DisplayName: "Ryan Ratajczak",
		Email:       "ryan@ratajczak.com",
		Provider:    "google",
	})
}
