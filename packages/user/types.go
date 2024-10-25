package main

type SignUpDetails struct {
	Email        string   `json:"email"`
	Provider     string   `json:"provider"`
	Password     string   `json:"password"`
	Username     string   `json:"username"`
	FirstName    string   `json:"firstName"`
	LastName     string   `json:"lastName"`
	Diet         string   `json:"diet"`
	Intolerances []string `json:"intolerances"`
}

type Profile struct {
	ID           string   `json:"id"`
	Username     string   `json:"username"`
	FirstName    string   `json:"firstName" db:"first_name"`
	LastName     string   `json:"lastName" db:"last_name"`
	Diet         string   `json:"diet"`
	Intolerances []string `json:"intolerances"`
}
