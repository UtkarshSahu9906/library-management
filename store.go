package main

// store holds all data in memory
var books = []Book{
	{
		ID:        "B001",
		Title:     "The Go Programming Language",
		Author:    "Alan Donovan",
		ISBN:      "978-0134190440",
		Available: true,
	},
	{
		ID:        "B002",
		Title:     "Clean Code",
		Author:    "Robert Martin",
		ISBN:      "978-0132350884",
		Available: true,
	},
	{
		ID:        "B003",
		Title:     "Design Patterns",
		Author:    "Gang of Four",
		ISBN:      "978-0201633610",
		Available: true,
	},
}

var members = []Member{
	{
		ID:    "M001",
		Name:  "Rahul Sharma",
		Email: "rahul@email.com",
		Phone: "9876543210",
	},
	{
		ID:    "M002",
		Name:  "Priya Singh",
		Email: "priya@email.com",
		Phone: "9123456780",
	},
}

var borrowRecords = []BorrowRecord{}