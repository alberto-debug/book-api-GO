package models

type Book struct {
	ID     string `json:"id"`
	NAME   string `json:"name"`
	AUTHOR string `json:"author"`
	YEAR   int    `json:"year"`
}
