package Library

import (
	"fmt"
)

type Book struct {
	ID         string
	Name       string
	Author     string
	isBorrowed bool
}

type Library struct {
	Books map[string]Book
}

func NewLibrary() *Library {
	return &Library{make(map[string]Book)}
}
func (l *Library) AddBook(book Book) {
	if _, exists := l.Books[book.ID]; exists {
		fmt.Println("Book already exists")
		return
	}

	l.Books[book.ID] = book
	fmt.Println("Book added")
}

func (l *Library) BorrowBook(id string) {
	book, exists := l.Books[id]
	if !exists {
		fmt.Println("There is no book with this id")
		return
	}

	if book.isBorrowed {
		fmt.Println("Book is already borrowed")
		return
	}
	book.isBorrowed = true
	fmt.Println("Book successfully borrowed")
}

func (l *Library) ReturnBook(id string) {
	book, exists := l.Books[id]
	if !exists {
		fmt.Println("There is no book with this id")
		return
	}

	if !book.isBorrowed {
		fmt.Println("Book is already returned")
		return
	}
	book.isBorrowed = false
	fmt.Println("Book returned borrowed")
}

func (l *Library) ListBooks() {
	for _, book := range l.Books {
		if book.isBorrowed {
			fmt.Printf("Book #%s: %s by %s. Not available\n", book.ID, book.Name, book.Author)
		} else {
			fmt.Printf("Book #%s: %s by %s. Available\n", book.ID, book.Name, book.Author)
		}
	}
}
