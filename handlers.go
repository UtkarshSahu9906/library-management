package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func borrowBook(w http.ResponseWriter, r *http.Request) {
	var request struct {
		BookID   string `json:"book_id"`
		MemberID string `json:"member_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	bookIndex := -1
	for i, b := range books {
		if b.ID == request.BookID {
			bookIndex = i
			break
		}
	}

	if bookIndex == -1 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	if !books[bookIndex].Available {
		http.Error(w, "Book is not available", http.StatusConflict)
		return
	}

	memberFound := false

	for _, m := range members {
		if m.ID == request.MemberID {
			memberFound = true
			break
		}
	}

	if !memberFound {
		http.Error(w, "Member not found", http.StatusNotFound)
		return
	}

	books[bookIndex].Available = false

	borrowDate := time.Now()
	dueDate := borrowDate.AddDate(0, 0, 14)
	record := BorrowRecord{
		ID:         fmt.Sprintf("BR%d", len(borrowRecords)+1),
		BookID:     request.BookID,
		MemberID:   request.MemberID,
		BorrowDate: borrowDate.Format("2006-01-02"),
		DueDate:    dueDate.Format("2006-01-02"),
		ReturnDate: "",
		Fine:       0,
		Returned:   false,
	}

	borrowRecords = append(borrowRecords, record)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(record)

}

// getAllBooks handles GET /books
func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// addBook handles POST /books
func addBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	books = append(books, book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}
