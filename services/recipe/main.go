package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	svc := NewRecipeService(os.Getenv("API_HOST"), os.Getenv("API_KEY"))
	api := NewApiServer(svc)
	log.Fatal(api.Start(":8080").Error())
}
