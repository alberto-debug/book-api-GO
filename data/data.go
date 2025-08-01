package data

import (
	"book-api/models"
	"errors"
)

var books = []models.Book{

	{ID: "1", Title: "The GO programming Language", Author: "Alberto Junior", Year: 2025},

	{ID: "2", Title: "CyberSecurity", Author: "Alberto Junior", Year: 2020},
}

func GetAllBooks() []models.Book {
	return books
}

func GetBookByID(id string) (models.Book, error) {
	for _, book := range books {

		if book.ID == id {

			return book, nil

		}

	}

	return models.Book{}, errors.New("Error not found")
}
