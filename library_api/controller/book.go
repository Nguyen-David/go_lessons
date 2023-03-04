package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"library_api/entity"
	"library_api/repository"
	"log"
	"net/http"
	"sort"
	"time"
)

type Book struct {
	Book_repository repository.BookRepository
}

func (b *Book) CreateBook(w http.ResponseWriter, r *http.Request) {
	book := entity.Book{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	err = json.Unmarshal(body, &book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	_, err = b.Book_repository.Create(book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	bl := entity.BookList{}
	err = json.Unmarshal(body, &bl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	fmt.Println("Successfully Create books")
	w.WriteHeader(http.StatusCreated)

}

func (b *Book) List(w http.ResponseWriter, r *http.Request) {
	books, err := b.Book_repository.List()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	sort.Sort(entity.SortedBooks(books))

	date := entity.BookTime(time.Now())

	bookList := entity.BookList{Books: books, Date: date}

	res, err := json.Marshal(bookList)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	// output as json file
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
