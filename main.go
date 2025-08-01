package main

import (
	"book-api/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize router
	r := mux.NewRouter()

	// Define endpoints
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")

	// Start server
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
