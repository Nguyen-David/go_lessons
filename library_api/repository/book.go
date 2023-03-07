package repository

import (
	"library_api/entity"
)

type BookRepository interface {
	Create(book entity.Book) (int, error)
	List() ([]entity.Book, error)
}

type bookRepository struct {
	books []entity.Book
}

func NewBookRepository() BookRepository {
	return &bookRepository{books: []entity.Book{
		{"Rage", "Stephen King", 1977},
		{"Philosopher's Stone", "J. K. Rowling", 1997},
		{"All Quiet on the Western Front", "Erich Maria Remarque", 1929},
	}}
}

func (r *bookRepository) Create(book entity.Book) (int, error) {
	r.books = append(r.books, book)

	return len(r.books), nil
}

func (r *bookRepository) List() ([]entity.Book, error) {
	return r.books, nil
}
