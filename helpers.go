package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// calculateDueDate returns the due date from borrow date
func calculateDueDate(borrowDate time.Time) time.Time {
	return borrowDate.AddDate(0, 0, 14)
}

// calculateFine returns the fine amount for a borrow record
func calculateFine(dueDate string, returnDate string) float64 {
	due, _ := time.Parse("2006-01-02", dueDate)
	returned, _ := time.Parse("2006-01-02", returnDate)

	diff := returned.Sub(due)
	days := int(diff.Hours() / 24)

	if days <= 0 {
		return 0
	}

	finePerDay := 5.0
	return float64(days) * finePerDay
}

// writeJSON sends a JSON response with a status code
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// writeError sends a plain text error response
func writeError(w http.ResponseWriter, status int, message string) {
	http.Error(w, message, status)
}

// validateBook checks if a book has all required fields
// bug: we only validate addBook, not borrowBook
func validateBook(book Book) string {
	if book.ID == "" {
		return "Book ID is required"
	}
	if book.Title == "" {
		return "Book title is required"
	}
	if book.Author == "" {
		return "Book author is required"
	}
	if book.ISBN == "" {
		return "Book ISBN is required"
	}
	return ""
}
// validateBorrowRequest checks if borrow request has required fields
func validateBorrowRequest(bookID string, memberID string) string {
	if bookID == "" {
		return "book_id is required"
	}
	if memberID == "" {
		return "member_id is required"
	}
	return ""
}

// validateReturnRequest checks if return request has required fields
func validateReturnRequest(borrowID string) string {
	if borrowID == "" {
		return "borrow_id is required"
	}
	return ""
}