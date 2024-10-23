package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	url := os.Getenv("RAPIDAPI_URL")

	svc := NewRecipeService(url)
	api := NewApiServer(svc)
	log.Fatal(api.Start(":3000").Error())
}
