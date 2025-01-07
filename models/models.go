package models

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Slug  string `json:"slug"`
	Date  string `json:"date"`
	Body  string `json:"body"`
}

type Posts struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Slug  string `json:"slug"`
	Date  string `json:"date"`
	Body  string `json:"body"`
}
