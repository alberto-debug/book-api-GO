package handlers

import (
	"encoding/json"
	"net/http"

	"book-api/data"

	"github.com/gorilla/mux"
)

// GetBooks returns all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books := data.GetAllBooks()
	json.NewEncoder(w).Encode(books)
}

// GetBook returns a single book by ID
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	book, err := data.GetBookByID(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}
