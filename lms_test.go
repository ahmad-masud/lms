package lms_test

import (
	"testing"

	"lms"
)

func setupLibrary() *lms.Library {
	return &lms.Library{
		Books:   make(map[int]*lms.Book),
		Members: make(map[int]*lms.Member),
	}
}

func TestAddBook(t *testing.T) {
	lib := setupLibrary()

	book := lms.Book{ID: 1, Title: "Go Basics", Author: "Alice", Copies: 3}
	lib.AddBook(book)

	if b, exists := lib.Books[1]; !exists {
		t.Fatal("Book not added")
	} else if b.Copies != 3 {
		t.Errorf("Expected 3 copies, got %d", b.Copies)
	}

	// Add more copies
	lib.AddBook(lms.Book{ID: 1, Title: "Go Basics", Author: "Alice", Copies: 2})
	if b := lib.Books[1]; b.Copies != 5 {
		t.Errorf("Expected 5 copies after adding more, got %d", b.Copies)
	}
}

func TestRegisterMember(t *testing.T) {
	lib := setupLibrary()

	member := lms.Member{ID: 1, Name: "John"}
	lib.RegisterMember(member)

	if _, exists := lib.Members[1]; !exists {
		t.Fatal("Member not registered")
	}
}

func TestBorrowAndReturnBook(t *testing.T) {
	lib := setupLibrary()
	lib.AddBook(lms.Book{ID: 1, Title: "Go Basics", Author: "Alice", Copies: 1})
	lib.RegisterMember(lms.Member{ID: 1, Name: "John"})

	// Borrow
	if err := lib.BorrowBook(1, 1); err != nil {
		t.Fatalf("BorrowBook failed: %v", err)
	}
	if lib.Books[1].Copies != 0 {
		t.Errorf("Expected 0 copies, got %d", lib.Books[1].Copies)
	}

	// Borrow again (should fail)
	if err := lib.BorrowBook(1, 1); err == nil {
		t.Fatal("Expected error when borrowing with 0 copies")
	}

	// Return
	if err := lib.ReturnBook(1, 1); err != nil {
		t.Fatalf("ReturnBook failed: %v", err)
	}
	if lib.Books[1].Copies != 1 {
		t.Errorf("Expected 1 copy after return, got %d", lib.Books[1].Copies)
	}
}

func TestListBooks(t *testing.T) {
	lib := setupLibrary()
	lib.AddBook(lms.Book{ID: 1, Title: "Go Basics", Author: "Alice", Copies: 1})
	lib.AddBook(lms.Book{ID: 2, Title: "Advanced Go", Author: "Bob", Copies: 2})

	titles := lib.ListBooks()
	expected := []string{"Advanced Go", "Go Basics"}

	for i, title := range titles {
		if title != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], title)
		}
	}
}

func TestListMemberBorrowed(t *testing.T) {
	lib := setupLibrary()
	lib.AddBook(lms.Book{ID: 1, Title: "Go Basics", Author: "Alice", Copies: 1})
	lib.RegisterMember(lms.Member{ID: 1, Name: "John"})

	lib.BorrowBook(1, 1)
	books, err := lib.ListMemberBorrowed(1)
	if err != nil {
		t.Fatalf("ListMemberBorrowed returned error: %v", err)
	}
	if len(books) != 1 || books[0] != "Go Basics" {
		t.Errorf("Expected ['Go Basics'], got %v", books)
	}
}

func TestSearchBooks(t *testing.T) {
	lib := setupLibrary()
	lib.AddBook(lms.Book{ID: 1, Title: "Go Basics", Author: "Alice", Copies: 1})
	lib.AddBook(lms.Book{ID: 2, Title: "Advanced Go", Author: "Bob", Copies: 2})
	lib.AddBook(lms.Book{ID: 3, Title: "Python 101", Author: "Carol", Copies: 1})

	results := lib.SearchBooks("go")
	if len(results) != 2 {
		t.Errorf("Expected 2 search results, got %d", len(results))
	}

	results = lib.SearchBooks("alice")
	if len(results) != 1 || results[0].ID != 1 {
		t.Errorf("Expected book ID 1 for 'alice', got %+v", results)
	}

	results = lib.SearchBooks("java")
	if len(results) != 0 {
		t.Errorf("Expected 0 results for 'java', got %+v", results)
	}
}
