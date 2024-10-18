package main

type Response struct {
	Recipes []Recipe `json:"results"`
}

type Recipe struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Image     string `json:"image"`
	ImageType string `json:"imageType"`
}
