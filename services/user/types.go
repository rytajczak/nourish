package main

import "user/repository"

type Profile struct {
	Diet     string `json:"diet"`
	Calories int    `json:"calories"`
	Protein  int    `json:"protein"`
	Carbs    int    `json:"carbs"`
	Fat      int    `json:"fat"`
}

type CreateUserRequest struct {
	Email        string   `json:"email"`
	Username     string   `json:"username"`
	Provider     string   `json:"provider"`
	Picture      string   `json:"picture"`
	Profile      Profile  `json:"profile"`
	Intolerances []string `json:"intolerances"`
}

type UserResponse struct {
	Profile         repository.GetUserProfileRow     `json:"profile"`
	Intolerances    []string                         `json:"intolerances"`
	SavedRecipes    []int                            `json:"savedRecipes"`
	SpoonCredential repository.GetUsernameAndHashRow `json:"spoonCredential"`
}
