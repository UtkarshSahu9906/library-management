package main

import (
	"encoding/json"
	"net/http"
)

// getAllBooks handles GET /books
// returns all books in the library as JSON
func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}