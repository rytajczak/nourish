package main

type RegisterDetails struct {
	Email        string   `json:"email"`
	Username     string   `json:"username"`
	FirstName    string   `json:"firstName"`
	LastName     string   `json:"lastName"`
	Diet         string   `json:"diet"`
	Intolerences []string `json:"intolerences"`
}
