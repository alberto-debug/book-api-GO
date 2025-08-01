package models

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"Title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}
