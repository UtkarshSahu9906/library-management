package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// getAllBooks handles GET /books
func getAllBooks(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, books)
}

// addBook handles POST /books
func addBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	books = append(books, book)
	writeJSON(w, http.StatusCreated, book)
}

// borrowBook handles POST /borrow
func borrowBook(w http.ResponseWriter, r *http.Request) {
	var request struct {
		BookID   string `json:"book_id"`
		MemberID string `json:"member_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
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

	if bookIndex == -1 {
		writeError(w, http.StatusNotFound, "Book not found")
		return
	}

	if !books[bookIndex].Available {
		writeError(w, http.StatusBadRequest, "Book is not available")
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

	if !memberFound {
		writeError(w, http.StatusNotFound, "Member not found")
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
	writeJSON(w, http.StatusCreated, record)
}

// getBorrowRecord handles GET /borrow/{id}
func getBorrowRecord(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/borrow/")

	if id == "" {
		writeError(w, http.StatusBadRequest, "Borrow record ID is required")
		return
	}

	for _, record := range borrowRecords {
		if record.ID == id {
			writeJSON(w, http.StatusOK, record)
			return
		}
	}

	writeError(w, http.StatusNotFound, "Borrow record not found")
}

// returnBook handles POST /return
func returnBook(w http.ResponseWriter, r *http.Request) {
	var request struct {
		BorrowID string `json:"borrow_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
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

	if recordIndex == -1 {
		writeError(w, http.StatusNotFound, "Borrow record not found")
		return
	}

	if borrowRecords[recordIndex].Returned {
		writeError(w, http.StatusBadRequest, "Book already returned")
		return
	}

	// mark book as available again
	for i, b := range books {
		if b.ID == borrowRecords[recordIndex].BookID {
			books[i].Available = true
			break
		}
	}

	// update borrow record
	returnDate := time.Now().Format("2006-01-02")
	borrowRecords[recordIndex].Returned = true
	borrowRecords[recordIndex].ReturnDate = returnDate
	borrowRecords[recordIndex].Fine = calculateFine(
		borrowRecords[recordIndex].DueDate,
		returnDate,
	)

	writeJSON(w, http.StatusOK, borrowRecords[recordIndex])
}

// getMemberFine handles GET /fine/{memberID}
func getMemberFine(w http.ResponseWriter, r *http.Request) {
	memberID := strings.TrimPrefix(r.URL.Path, "/fine/")

	if memberID == "" {
		writeError(w, http.StatusBadRequest, "Member ID is required")
		return
	}

	memberFound := false
	for _, m := range members {
		if m.ID == memberID {
			memberFound = true
			break
		}
	}

	if !memberFound {
		writeError(w, http.StatusNotFound, "Member not found")
		return
	}

	totalFine := 0.0
	var memberRecords []BorrowRecord
	for _, record := range borrowRecords {
		if record.MemberID == memberID {
			totalFine += record.Fine
			memberRecords = append(memberRecords, record)
		}
	}

	response := struct {
		MemberID  string         `json:"member_id"`
		TotalFine float64        `json:"total_fine"`
		Records   []BorrowRecord `json:"records"`
	}{
		MemberID:  memberID,
		TotalFine: totalFine,
		Records:   memberRecords,
	}

	writeJSON(w, http.StatusOK, response)
}