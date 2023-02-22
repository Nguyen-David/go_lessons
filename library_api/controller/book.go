package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"library_api/entity"
	"library_api/repository"
	"net/http"
	"sort"
)

type Book struct {
	Book_repository repository.BookRepository
}

func (b *Book) CreateBooks(w http.ResponseWriter, r *http.Request) {
	book := entity.Book{}

	body, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &book)

	_, err = b.Book_repository.Create(book)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully Create books")
	w.WriteHeader(http.StatusCreated)

}

func (b *Book) List(w http.ResponseWriter, r *http.Request) {
	books, err := b.Book_repository.List()
	if err != nil {
		panic(err)
	}

	sort.Sort(entity.ByYear(books))

	res, err := json.Marshal(books) 
	if err != nil {
		panic(err)
	}

	// output as json file
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}
