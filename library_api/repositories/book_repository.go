package repositories

import (
	"library_api/entities"
)

type BookRepository interface {
	Create(books []entities.Book) (int, error)
    List() ([]entities.Book, error)
}

type bookRepository struct {
	books []entities.Book
}

func NewBookRepository() BookRepository {
	return &bookRepository{}
}

func (r *bookRepository) Create(books []entities.Book) (int, error) {
	r.books = books

	return len(r.books), nil
}

func (r *bookRepository) List() ([]entities.Book, error) {
	return r.books, nil
}