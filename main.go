package main

import (
	"book-api/handlers"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
}
