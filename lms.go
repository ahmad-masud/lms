package lms

import (
	"errors"
	"sort"
	"fmt"
	"strings"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Copies int
}

type Member struct {
	ID       int
	Name     string
	Borrowed []int
}

type Library struct {
	Books   map[int]*Book
	Members map[int]*Member
}

type Stringer interface {
	String() string
}

func (library *Library) AddBook(book Book) {
	b, exists := library.Books[book.ID]

	if exists {
		b.Copies += book.Copies
	} else {
		newBook := book
		library.Books[book.ID] = &newBook
	}
}

func (library *Library) RegisterMember(member Member) {
	newMember := member
	library.Members[member.ID] = &newMember
}

func (library *Library) BorrowBook(memberID int, bookID int) error {
	m, mExists := library.Members[memberID]
	b, bExists := library.Books[bookID]

	if mExists && bExists {
		if b.Copies > 0 {
			b.Copies--
			m.Borrowed = append(m.Borrowed, bookID)
		} else {
			return errors.New("No more copies")
		}
	} else {
		return errors.New("Member or Book does not exist")
	}

	return nil
}

func (library *Library) ReturnBook(memberID int, bookID int) error {
	m, mExists := library.Members[memberID]
	b, bExists := library.Books[bookID]

	if mExists && bExists {
		b.Copies++

		for i, bi := range m.Borrowed {
			if bi == bookID {
				m.Borrowed = append(m.Borrowed[:i], m.Borrowed[i+1:]...)
				break
			}
		}
	} else {
		return errors.New("Member or Book does not exist")
	}

	return nil
}

func (library *Library) ListBooks() []string {
	books := []string{}

	for _, book := range library.Books {
		books = append(books, book.Title)
	}

	sort.Strings(books)

	return books
}

func (library *Library) ListMemberBorrowed(memberID int) ([]string, error) {
	member, exists := library.Members[memberID]
	books := []string{}

	if exists {
		for _, bookID := range member.Borrowed {
			books = append(books, library.Books[bookID].Title)
		}
	} else {
		return nil, errors.New("Member does not exist")
	}

	return books, nil
}

func (book Book) String() string {
	return fmt.Sprintf("%s by %s (%d copies available)", book.Title, book.Author, book.Copies)
}

func (member Member) String() string {
	return fmt.Sprintf("%s borrowed %d books", member.Name, len(member.Borrowed))
}

func (library *Library) SearchBooks(keyword string) []Book {
	books := []Book{}
	keyword = strings.ToLower(keyword)

	for _, book := range library.Books {
		title := strings.ToLower(book.Title)
		author := strings.ToLower(book.Author)
		
		if strings.Contains(title, keyword) || strings.Contains(author, keyword) {
			books = append(books, *book)
		}
	}

	return books
}