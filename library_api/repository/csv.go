package repository

import (
	"library_api/entity"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type csvRepository struct {
	fileConnection *os.File
}

func NewCsvRepository(fileConnection *os.File) BookRepository {
	return &csvRepository{fileConnection: fileConnection}
}

func (r *csvRepository) Create(book entity.Book) (int, error) {
	books, err := r.List()
	if err != nil { // Load clients from file
		log.Println(err)
		return 0, err
	}

	if _, err := r.fileConnection.Seek(0, 0); err != nil {
		return 0, err
	}

	books = append(books, book)

	if err := gocsv.MarshalFile(books, r.fileConnection); err != nil { // Load clients from file
		log.Println(err)
		return 0, err
	}

	return len(books), nil
}

func (r *csvRepository) List() ([]entity.Book, error) {
	if _, err := r.fileConnection.Seek(0, 0); err != nil {
		return nil, err
	}

	var books []entity.Book
	println(r.fileConnection)
	if err := gocsv.UnmarshalFile(r.fileConnection, &books); err != nil { // Load clients from file
		log.Println(err)
		return nil, err
	}

	return books, nil
}
