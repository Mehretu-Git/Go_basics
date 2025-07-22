📚 Library Management System Documentation
📌 Overview
This is a simple console-based library management system implemented in Go (Golang). It demonstrates the use of:

Structs

Interfaces

Slices

Maps

Console interaction

📁 Folder Structure
go
Copy
Edit
library_management/
├── main.go
├── controllers/
│   └── library_controller.go
├── models/
│   ├── book.go
│   └── member.go
├── services/
│   └── library_service.go
├── docs/
│   └── documentation.md
└── go.mod
🧩 Structs
Book

ID (int)

Title (string)

Author (string)

Status (string) — "Available" or "Borrowed"

Member

ID (int)

Name (string)

BorrowedBooks ([]Book)

🔌 Interface
LibraryManager

AddBook(book Book)

RemoveBook(bookID int)

BorrowBook(bookID int, memberID int) string

ReturnBook(bookID int, memberID int) string

ListAvailableBooks() []Book

ListBorrowedBooks(memberID int) []Book

🏗️ Features
Add a new book

Remove a book by ID

Borrow a book (if available)

Return a borrowed book

List available books

List books borrowed by a member

🖥️ Console Interaction
You can run the program using:

bash
Copy
Edit
go run main.go
You will then be prompted with a console interface to choose operations.

✅ Usage Example
markdown
Copy
Edit
1. Add Book
2. Remove Book
3. Borrow Book
4. Return Book
5. List Available Books
6. List Borrowed Books by Member
7. Exit
🛠️ Dependencies
This project uses only the Go standard library. No external packages are required.

⚠️ Notes
No use of the errors package — all error messages are handled with string return values.

Book and member data are stored in memory (not persistent).

