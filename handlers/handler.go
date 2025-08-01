package handlers

import (
	"book-api/data"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books := data.GetAllBooks()
	json.NewEncoder(w).Encode(books)

}

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
