package Library

import (
	"testing"
)

func TestLibrary(t *testing.T) {
	// Create library
	library := NewLibrary()
	library.AddBook(Book{ID: "1", Name: "title1", Author: "author1"})
	library.AddBook(Book{ID: "2", Name: "title2", Author: "author2"})
	library.AddBook(Book{ID: "3", Name: "title3", Author: "author3"})
	library.AddBook(Book{ID: "1", Name: "title4", Author: "author4"})
	library.ListBooks()

	library.BorrowBook("4")
	library.ListBooks()
	library.ReturnBook("1")
	library.ListBooks()

}
