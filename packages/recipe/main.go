package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	url := os.Getenv("RAPIDAPI_URL")

	svc := NewRecipeService(url)
	api := NewApiServer(svc)
	api.Start(":8080")
}
