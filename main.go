package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Library Management System starting...")

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to Library Management System")
	})

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

	mux.HandleFunc("/borrow", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			borrowBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	port := ":8080"
	log.Println("Server starting on port", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}