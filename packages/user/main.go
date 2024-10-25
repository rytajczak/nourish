package main

import "log"

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal("couldn't connect to db")
	}

	if err := store.db.Ping(); err != nil {
		log.Fatalf("couldn't establish connection to db: %s", err.Error())
	}
	log.Println("successfully connected to database")

	api := NewApiServer(store)
	log.Fatal(api.Start(":8082").Error())
}
