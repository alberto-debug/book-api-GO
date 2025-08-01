package handlers

import (
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	books := data.
		json.NewEncoder(w).Encode(books)

}
