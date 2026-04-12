package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)




// getMemberFine handles GET /fine/{memberID}
// returns total fine for a member
func getMemberFine(w http.ResponseWriter, r *http.Request) {
	memberID := strings.TrimPrefix(r.URL.Path, "/fine/")

	if memberID == "" {
		http.Error(w, "Member ID is required", http.StatusBadRequest)
		return
	}

	// check member exists
	memberFound := false
	for _, m := range members {
		if m.ID == memberID {
			memberFound = true
			break
		}
	}

	if !memberFound {
		http.Error(w, "Member not found", http.StatusNotFound)
		return
	}

	// add up all fines for this member
	totalFine := 0.0
	var memberRecords []BorrowRecord
	for _, record := range borrowRecords {
		if record.MemberID == memberID {
			totalFine += record.Fine
			memberRecords = append(memberRecords, record)
		}
	}

	// build response
	response := struct {
		MemberID    string         `json:"member_id"`
		TotalFine   float64        `json:"total_fine"`
		Records     []BorrowRecord `json:"records"`
	}{
		MemberID:  memberID,
		TotalFine: totalFine,
		Records:   memberRecords,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
// returnBook handles POST /return
// marks a book as returned
func returnBook(w http.ResponseWriter, r *http.Request) {
	var request struct {
		BorrowID string `json:"borrow_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// find the borrow record
	recordIndex := -1
	for i, record := range borrowRecords {
		if record.ID == request.BorrowID {
			recordIndex = i
			break
		}
	}

	// record not found
	if recordIndex == -1 {
		http.Error(w, "Borrow record not found", http.StatusNotFound)
		return
	}

	// already returned
	if borrowRecords[recordIndex].Returned {
		http.Error(w, "Book already returned", http.StatusBadRequest)
		return
	}

	// mark book as available again
	for i, b := range books {
		if b.ID == borrowRecords[recordIndex].BookID {
			books[i].Available = true
			break
		}
	}


returnDate := time.Now().Format("2006-01-02")
borrowRecords[recordIndex].Returned = true
borrowRecords[recordIndex].ReturnDate = returnDate
borrowRecords[recordIndex].Fine = calculateFine(
    borrowRecords[recordIndex].DueDate,
    returnDate,
)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(borrowRecords[recordIndex])
}

func getBorrowRecord(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/borrow/")

	if id == "" {
		http.Error(w, "Borrow ID is required", http.StatusBadRequest)
		return
	}
	for _, record := range borrowRecords {
		if record.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(record)
			return
		}

	}

	http.Error(w, "Borrow record not found", http.StatusNotFound)
}

// borrowBook handles POST /borrow
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

	// find the book
	bookIndex := -1
	for i, b := range books {
		if b.ID == request.BookID {
			bookIndex = i
			break
		}
	}

	// book not found
	if bookIndex == -1 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// book already borrowed
	if !books[bookIndex].Available {
		http.Error(w, "Book is not available", http.StatusBadRequest)
		return
	}

	// find the member
	memberFound := false
	for _, m := range members {
		if m.ID == request.MemberID {
			memberFound = true
			break
		}
	}

	// member not found
	if !memberFound {
		http.Error(w, "Member not found", http.StatusNotFound)
		return
	}

	// mark book as unavailable
	books[bookIndex].Available = false

	// create borrow record
	borrowDate := time.Now()
	dueDate := calculateDueDate(borrowDate)

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
