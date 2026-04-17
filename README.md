markdown# 
Library Management System

A REST API built with Go to manage a library's books, members, borrowing, returns, and fine calculations.

Built as a learning project to understand Go REST API development — including real bug fixes and iterative commits.

---

## Tech Stack

- **Language:** Go
- **Router:** net/http (standard library)
- **Storage:** In-memory (slice-based)
- **No external dependencies**

---

## Project Structure
library-management/
├── main.go        → server setup and routes
├── models.go      → data structs (Book, Member, BorrowRecord)
├── store.go       → in-memory data store
├── handlers.go    → HTTP handler functions
├── helpers.go     → utility functions (fine, validation, response)
└── go.mod         → Go module file

---

## Getting Started

### Prerequisites
- Go 1.21 or higher installed

### Run the server

```bash
git clone https://github.com/YOUR_USERNAME/library-management.git
cd library-management
go run .
```

Server starts on `http://localhost:8080`

---

## API Endpoints

### Books

#### Get all books
GET /books

Response:
```json
[
  {
    "id": "B001",
    "title": "The Go Programming Language",
    "author": "Alan Donovan",
    "isbn": "978-0134190440",
    "available": true
  }
]
```

---

#### Add a new book
POST /books

Request body:
```json
{
  "id": "B004",
  "title": "Go in Action",
  "author": "Kennedy",
  "isbn": "978-1617291784",
  "available": true
}
```

Response: `201 Created`
```json
{
  "id": "B004",
  "title": "Go in Action",
  "author": "Kennedy",
  "isbn": "978-1617291784",
  "available": true
}
```

---

### Borrowing

#### Borrow a book
POST /borrow

Request body:
```json
{
  "book_id": "B001",
  "member_id": "M001"
}
```

Response: `201 Created`
```json
{
  "id": "BR1",
  "book_id": "B001",
  "member_id": "M001",
  "borrow_date": "2026-04-09",
  "due_date": "2026-04-23",
  "return_date": "",
  "fine": 0,
  "returned": false
}
```

---

#### Check borrow record status
GET /borrow/{id}

Example: `GET /borrow/BR1`

Response: `200 OK`
```json
{
  "id": "BR1",
  "book_id": "B001",
  "member_id": "M001",
  "borrow_date": "2026-04-09",
  "due_date": "2026-04-23",
  "return_date": "",
  "fine": 0,
  "returned": false
}
```

---

### Returns

#### Return a book
POST /return

Request body:
```json
{
  "borrow_id": "BR1"
}
```

Response: `200 OK`
```json
{
  "id": "BR1",
  "book_id": "B001",
  "member_id": "M001",
  "borrow_date": "2026-04-09",
  "due_date": "2026-04-23",
  "return_date": "2026-04-09",
  "fine": 0,
  "returned": true
}
```

---

### Fines

#### Get total fine for a member
GET /fine/{memberID}

Example: `GET /fine/M001`

Response: `200 OK`
```json
{
  "member_id": "M001",
  "total_fine": 35,
  "records": [
    {
      "id": "BR1",
      "book_id": "B001",
      "member_id": "M001",
      "borrow_date": "2026-04-09",
      "due_date": "2026-04-23",
      "return_date": "2026-04-30",
      "fine": 35,
      "returned": true
    }
  ]
}
```

Fine is calculated at **5 rupees per day** after the due date.
Books can be borrowed for **14 days** before the due date.
No fine if returned on time or early.

---

## Validation Rules

### POST /books
| Field | Rule |
|---|---|
| id | required |
| title | required |
| author | required |
| isbn | required |

### POST /borrow
| Field | Rule |
|---|---|
| book_id | required |
| member_id | required |

### POST /return
| Field | Rule |
|---|---|
| borrow_id | required |

---

## Error Responses

All errors return plain text with appropriate status code:
400 Bad Request   → invalid input or validation failure
404 Not Found     → resource does not exist
405 Not Allowed   → wrong HTTP method used

Examples:
Book ID is required
Book not found
Book is not available
Member not found
Borrow record not found
Book already returned
book_id is required
member_id is required
borrow_id is required

---

## Fine Calculation Logic
Borrow date  → day book is taken
Due date     → borrow date + 14 days
Return date  → day book is brought back
If return date <= due date → fine = 0
If return date >  due date → fine = (days late) × 5 rupees

Examples:
Borrowed April 9  → due April 23
Returned April 23 → 0 days late  → fine = 0
Returned April 24 → 1 day late   → fine = 5
Returned April 30 → 7 days late  → fine = 35
Returned May  23  → 30 days late → fine = 150

---

## Sample Data

The system starts with this data preloaded:

### Books
| ID | Title | Author | ISBN |
|---|---|---|---|
| B001 | The Go Programming Language | Alan Donovan | 978-0134190440 |
| B002 | Clean Code | Robert Martin | 978-0132350884 |
| B003 | Design Patterns | Gang of Four | 978-0201633610 |

### Members
| ID | Name | Email | Phone |
|---|---|---|---|
| M001 | Rahul Sharma | rahul@email.com | 9876543210 |
| M002 | Priya Singh | priya@email.com | 9123456780 |

---

## Status Codes

| Code | Meaning | When |
|---|---|---|
| 200 | OK | successful GET or return |
| 201 | Created | successful POST that creates data |
| 400 | Bad Request | invalid input or validation failure |
| 404 | Not Found | resource does not exist |
| 405 | Method Not Allowed | wrong HTTP method |

---

## Testing the API on Windows PowerShell

### GET request
```powershell
curl.exe http://localhost:8080/books
```

### POST request
```powershell
$body = '{"book_id":"B001","member_id":"M001"}'
Invoke-WebRequest -Uri http://localhost:8080/borrow -Method POST -ContentType "application/json" -Body $body -UseBasicParsing
```

---



## What I Learned

- Building a REST API in Go using only the standard library
- Structuring a Go project across multiple files
- In-memory data storage using slices
- HTTP handlers and routing with ServeMux
- JSON encoding and decoding
- Input validation and error handling
- Due date and fine calculation logic
- Iterative development with meaningful git commits
- Debugging and fixing real bugs in a running server

---

## Future Improvements

- [ ] Add a real database (PostgreSQL or SQLite)
- [ ] Add member registration endpoints
- [ ] Add JWT authentication
- [ ] Add unit tests
- [ ] Add pagination for book listing
- [ ] Add search books by title or author
- [ ] Add book categories
- [ ] Deploy to a cloud server
