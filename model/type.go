package model

type Category struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"imageUrl"`
}

type Recipe struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"imageUrl"`
}
