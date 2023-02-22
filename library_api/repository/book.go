package repository

import (
	"library_api/entity"
)

type BookRepository interface {
	Create(books []entity.Book) (int, error)
    List() ([]entity.Book, error)
}

type bookRepository struct {
	books []entity.Book
}

func NewBookRepository() BookRepository {
	return &bookRepository{}
}

func (r *bookRepository) Create(books []entity.Book) (int, error) {
	r.books = books

	return len(r.books), nil
}

func (r *bookRepository) List() ([]entity.Book, error) {
	return r.books, nil
}