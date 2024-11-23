package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	logger := log.New(os.Stderr)
	svc := NewUserService(os.Getenv("API_HOST"), os.Getenv("API_KEY"))
	svc = NewLoggingService(svc, logger)
	api := NewApiServer(svc)
	api.Start(":8081")
}
