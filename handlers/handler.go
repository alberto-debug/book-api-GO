package handlers

import (
	"book-api/data"
	"encoding/json"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	books := data.GetAllBooks()
	json.NewEncoder(w).Encode(books)

}
