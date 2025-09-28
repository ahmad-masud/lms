# Go Coding Assignment: Library Management System

## Overview
You are tasked with building a small in-memory library system in Go. The system should manage **books**, **members**, and **loans**. This assignment will test your understanding of arrays, slices, maps, structs, functions, methods, pointers, and interfaces.

---

## Part 1: Define Structs and Types

1. **Book struct**
   - Fields: `ID int`, `Title string`, `Author string`, `Copies int`
2. **Member struct**
   - Fields: `ID int`, `Name string`, `Borrowed []int` (IDs of borrowed books)
3. **Library struct**
   - Fields: 
     - `Books map[int]Book` – stores all books by ID
     - `Members map[int]Member` – stores all members by ID

---

## Part 2: Functions

Implement the following functions:

1. `AddBook(library *Library, book Book)`  
   - Adds a book to the library.  
   - If the book already exists, increase its `Copies`.

2. `RegisterMember(library *Library, member Member)`  
   - Adds a member to the library.  

3. `BorrowBook(library *Library, memberID int, bookID int) error`  
   - Member borrows a book.  
   - Reduce the book’s `Copies` by 1.  
   - Add bookID to member’s `Borrowed` slice.  
   - Return an error if book is unavailable or member does not exist.

4. `ReturnBook(library *Library, memberID int, bookID int) error`  
   - Member returns a book.  
   - Increase the book’s `Copies` by 1.  
   - Remove bookID from member’s `Borrowed` slice.  
   - Return an error if book or member not found.

---

## Part 3: Slices and Arrays

1. Write a function `ListBooks(library *Library) []string`  
   - Returns a slice of strings with book titles.  
   - Use slice operations to sort the titles alphabetically.  

2. Write a function `ListMemberBorrowed(library *Library, memberID int) ([]string, error)`  
   - Returns a slice of book titles that the member has borrowed.  

---

## Part 4: Maps

- All books and members should be stored in **maps**.  
- Accessing books/members must use the `ok` idiom to safely handle missing keys.

---

## Part 5: Methods

Convert some functions into **methods**:

1. `func (l *Library) AddBook(book Book)`  
2. `func (l *Library) RegisterMember(member Member)`  
3. `func (l *Library) BorrowBook(memberID int, bookID int) error`  
4. `func (l *Library) ReturnBook(memberID int, bookID int) error`  

---

## Part 6: Interfaces

1. Create an interface `Stringer` with a method `String() string`  
2. Implement `String()` for:
   - `Book` → returns `"Title by Author (Copies copies available)"`  
   - `Member` → returns `"Name borrowed X books"`  

---

## Part 7: Bonus Features (Optional)

Implement a **search function**:  
   ```go
   func (l *Library) SearchBooks(keyword string) []Book
   ```
   - Return books where `Title` or `Author` contains the keyword.

---

## Part 8: Example Usage
```go
func main() {
    library := Library{
        Books:   make(map[int]Book),
        Members: make(map[int]Member),
    }

    library.AddBook(Book{ID: 1, Title: "Go Programming", Author: "Alice", Copies: 3})
    library.AddBook(Book{ID: 2, Title: "Data Structures", Author: "Bob", Copies: 2})

    library.RegisterMember(Member{ID: 1, Name: "John"})
    library.RegisterMember(Member{ID: 2, Name: "Jane"})

    library.BorrowBook(1, 1)
    fmt.Println(library.ListMemberBorrowed(1)) // ["Go Programming"]
    fmt.Println(library.Books[1].Copies)       // 2

    library.ReturnBook(1, 1)
    fmt.Println(library.ListMemberBorrowed(1)) // []
    fmt.Println(library.Books[1].Copies)       // 3

    // Print all books
    for _, b := range library.Books {
        fmt.Println(b.String())
    }
}
```

---

## Learning Goals
This assignment tests the following Go concepts:

- Arrays, slices, and slice operations
- Maps and safe access
- Structs and pointers
- Methods and receivers
- Interfaces
- Functions (including variadic and returning multiple values)
- Error handling
- Sorting slices
- Working with zero values
