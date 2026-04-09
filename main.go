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

	port := ":8080"

	log.Println("Server starting on port", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}