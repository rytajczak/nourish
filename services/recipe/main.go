package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	host := os.Getenv("API_HOST")
	key := os.Getenv("API_KEY")

	svc := NewRecipeService(host, key)
	api := NewApiServer(svc)
	log.Fatal(api.Start(":8080").Error())
}
