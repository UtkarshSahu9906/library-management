package main

// Book represents a book in the library
type Book struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	ISBN      string `json:"isbn"`
	Available bool   `json:"available"`
}

// Member represents a library member
type Member struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// BorrowRecord tracks which member borrowed which book
type BorrowRecord struct {
	ID         string  `json:"id"`
	BookID     string  `json:"book_id"`
	MemberID   string  `json:"member_id"`
	BorrowDate string  `json:"borrow_date"`
	DueDate    string  `json:"due_date"`
	ReturnDate string  `json:"return_date"`
	Fine       float64 `json:"fine"`
	Returned   bool    `json:"returned"`
}