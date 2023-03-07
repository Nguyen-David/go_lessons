package repository

import (
	"library_api/config"
	"library_api/entity"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type csvRepository struct {
	books []entity.Book
}

func NewCsvRepository() BookRepository {
	return &csvRepository{}
}

func (r *csvRepository) Create(book entity.Book) (int, error) {
	r.books = append(r.books, book)

	booksFile, err := os.OpenFile(config.ConfigGlobal.RepoFilePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Println(err)
		return 0 ,err
	}
	defer booksFile.Close()

	if err := gocsv.MarshalFile(r.books, booksFile); err != nil { // Load clients from file
		log.Println(err)
		return 0 ,err
	}

	return len(r.books), nil
}

func (r *csvRepository) List() ([]entity.Book, error) {
	booksFile, err := os.OpenFile(config.ConfigGlobal.RepoFilePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Println(err)
		return r.books, err
	}
	defer booksFile.Close()

	if err := gocsv.UnmarshalFile(booksFile, &r.books); err != nil { // Load clients from file
		log.Println(err)
		return r.books, err
	}

	return r.books, nil
}