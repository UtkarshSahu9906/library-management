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

	log.Println("Server starting on port 9090...")
	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}