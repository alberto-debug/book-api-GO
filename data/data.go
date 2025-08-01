package data

import (
	"errors"

	"book-api/models"
)

// In-memory book store
var books = []models.Book{
	{ID: "1", Title: "The Go Programming Language", Author: "Alan Donovan", Year: 2015},
	{ID: "2", Title: "Learning Go", Author: "Jon Bodner", Year: 2021},
}

// GetAllBooks returns all books
func GetAllBooks() []models.Book {
	return books
}

// GetBookByID returns a book by its ID
func GetBookByID(id string) (models.Book, error) {
	for _, book := range books {
		if book.ID == id {
			return book, nil
		}
	}
	return models.Book{}, errors.New("book not found")
}
