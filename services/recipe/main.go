package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	svc := NewRecipeService(os.Getenv("API_HOST"), os.Getenv("API_KEY"))
	api := NewApiServer(svc)
	api.Start(":8082")
}
