package data

import (
	"book-api/models"
	"errors"
)

var books = []models.Book{
	{ID: "1", NAME: "CyberSecurity Fundamentals", AUTHOR: "Alberto Junior", YEAR: 2025},
	{ID: "2", NAME: "Programming Language", AUTHOR: "Alberto Manuel", YEAR: 2020},
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

	return models.Book{}, errors.New("Book not found with id: " + id)

}
