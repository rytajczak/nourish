package main

func main() {
	svc := NewRecipeService()
	apiServer := NewApiServer(svc)
	apiServer.Start(":8080")
}
