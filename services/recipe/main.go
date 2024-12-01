package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	cache := NewCache("recipe-cache:6379")

	svc := NewRecipeService(os.Getenv("API_HOST"), os.Getenv("API_KEY"), cache)
	svc = NewLoggingService(svc)

	api := NewApiServer(svc)
	api.Start(":8082")
}
