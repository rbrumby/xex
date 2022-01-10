package xex

import (
	"errors"
)

type Library struct {
	Address Address
	Books   []*Book
}

func (l Library) GetAddress() Address {
	return l.Address
}

func (l Library) Authors() map[int]string {
	authors := make(map[int]string)
	for _, b := range l.Books {
		authors[b.Author.Id] = b.Author.Name
	}
	return authors
}

func (l Library) GetBooks() []*Book {
	return l.Books
}

func (l Library) Book(title string) (*Book, error) {
	for _, b := range l.Books {
		if b.Title == title {
			return b, nil
		}
	}
	return nil, errors.New("Book not found")
}

type Address struct {
	Building string
	Street   string
	City     string
}

type Book struct {
	Title           string
	Author          *Author
	PublicationYear int
	Price           float32
}

type Author struct {
	Id   int
	Name string
}

func (a *Author) Books(l Library) ([]*Book, error) {
	books := make([]*Book, 0)
	for _, b := range l.Books {
		if b.Author == a {
			books = append(books, b)
		}
	}
	if len(books) == 0 {
		return nil, errors.New("No books for author")
	}
	return books, nil
}

var testLib = Library{
	Address: Address{
		Building: "123",
		Street:   "New Street",
		City:     "London",
	},
	Books: []*Book{
		{
			Title:           "Sense & Sensibility",
			PublicationYear: 1811,
			Price:           4.99,
			Author: &Author{
				Id:   1,
				Name: "Jane Austen",
			},
		},
		{
			Title:           "Pride & Prejudice",
			PublicationYear: 1813,
			Price:           6.99,
			Author: &Author{
				Id:   1,
				Name: "Jane Austen",
			},
		},
		{
			Title:           "1984",
			PublicationYear: 1949,
			Price:           9.99,
			Author: &Author{
				Id:   2,
				Name: "George Orwell",
			},
		},
		{
			Title:           "Animal Farm",
			PublicationYear: 1945,
			Price:           8.99,
			Author: &Author{
				Id:   2,
				Name: "George Orwell",
			},
		},
		{
			Title:           "The Lion, the With & the Wardrobe",
			PublicationYear: 1950,
			Price:           5.49,
			Author: &Author{
				Id:   3,
				Name: "C.S. Lewis",
			},
		},
	},
}
