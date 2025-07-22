package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"library_management/models"
	"library_management/services"
)

var reader = bufio.NewReader(os.Stdin)
var library = services.NewLibrary()

func StartConsole() {
	for {
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Exit")
		fmt.Print("Choose an option: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			addBook()
		case "2":
			removeBook()
		case "3":
			borrowBook()
		case "4":
			returnBook()
		case "5":
			listAvailableBooks()
		case "6":
			listBorrowedBooks()
		case "7":
			return
		}
	}
}

func addBook() {
	fmt.Print("Enter Book ID: ")
	id := readInt()
	fmt.Print("Enter Title: ")
	title := readLine()
	fmt.Print("Enter Author: ")
	author := readLine()
	book := models.Book{ID: id, Title: title, Author: author, Status: "Available"}
	library.AddBook(book)
}

func removeBook() {
	fmt.Print("Enter Book ID to remove: ")
	id := readInt()
	library.RemoveBook(id)
}

func borrowBook() {
	fmt.Print("Enter Book ID to borrow: ")
	bookID := readInt()
	fmt.Print("Enter Member ID: ")
	memberID := readInt()
	msg := library.BorrowBook(bookID, memberID)
	if msg != "" {
		fmt.Println("Error:", msg)
	}
}

func returnBook() {
	fmt.Print("Enter Book ID to return: ")
	bookID := readInt()
	fmt.Print("Enter Member ID: ")
	memberID := readInt()
	msg := library.ReturnBook(bookID, memberID)
	if msg != "" {
		fmt.Println("Error:", msg)
	}
}

func listAvailableBooks() {
	books := library.ListAvailableBooks()
	for _, book := range books {
		fmt.Println(book)
	}
}

func listBorrowedBooks() {
	fmt.Print("Enter Member ID: ")
	memberID := readInt()
	books := library.ListBorrowedBooks(memberID)
	for _, book := range books {
		fmt.Println(book)
	}
}

func readLine() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func readInt() int {
	input := readLine()
	val, _ := strconv.Atoi(input)
	return val
}
