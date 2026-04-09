package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Library Management System starting...")

	mux := http.NewServeMux()

	mux.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        getAllBooks(w, r)
    case http.MethodPost:
        addBook(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
})

	//mux.HandleFunc("/books", getAllBooks)

	port := ":8080"
	log.Println("Server starting on port", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}