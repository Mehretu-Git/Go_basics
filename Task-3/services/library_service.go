package services

import "library_management/models"

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) string
	ReturnBook(bookID int, memberID int) string
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	books   map[int]models.Book
	members map[int]*models.Member
}

func NewLibrary() LibraryManager {
	return &Library{
		books:   make(map[int]models.Book),
		members: make(map[int]*models.Member),
	}
}

func (l *Library) AddBook(book models.Book) {
	l.books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) {
	delete(l.books, bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) string {
	book, ok := l.books[bookID]
	if !ok {
		return "book not found"
	}
	if book.Status != "Available" {
		return "book already borrowed"
	}
	member, ok := l.members[memberID]
	if !ok {
		member = &models.Member{ID: memberID}
		l.members[memberID] = member
	}
	book.Status = "Borrowed"
	l.books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	return ""
}

func (l *Library) ReturnBook(bookID int, memberID int) string {
	book, ok := l.books[bookID]
	if !ok {
		return "book not found"
	}
	member, ok := l.members[memberID]
	if !ok {
		return "member not found"
	}
	book.Status = "Available"
	l.books[bookID] = book
	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}
	return ""
}

func (l *Library) ListAvailableBooks() []models.Book {
	var available []models.Book
	for _, book := range l.books {
		if book.Status == "Available" {
			available = append(available, book)
		}
	}
	return available
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, ok := l.members[memberID]
	if !ok {
		return nil
	}
	return member.BorrowedBooks
}
