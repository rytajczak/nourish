package main

type SignUpDetails struct {
	Email        string   `json:"email"`
	Provider     string   `json:"provider"`
	Password     string   `json:"password"`
	Username     string   `json:"username"`
	FirstName    string   `json:"firstName"`
	LastName     string   `json:"lastName"`
	Diet         string   `json:"diet"`
	Intolerences []string `json:"intolerences"`
}

type Profile struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Diet      string `json:"diet"`
}
