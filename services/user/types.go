package main

import (
	"fmt"
	"strings"
)

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
	ID           string      `json:"id"`
	Username     string      `json:"username"`
	FirstName    string      `json:"firstName" db:"first_name"`
	LastName     string      `json:"lastName" db:"last_name"`
	Diet         string      `json:"diet"`
	Intolerances StringArray `json:"intolerances"`
}

// wow this sucks! (https://github.com/jmoiron/sqlx/issues/578)
type StringArray []string

func (s *StringArray) Scan(value interface{}) error {
	if value == nil {
		*s = StringArray{}
		return nil
	}

	// Ensure the value is of type []byte (string in PostgreSQL)
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan StringArray: %v", value)
	}

	// Remove the surrounding braces and split by comma
	str := string(bytes)
	str = str[1 : len(str)-1] // remove { and }
	*s = StringArray(strings.Split(str, ","))

	return nil
}
