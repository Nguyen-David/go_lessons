package controllers

import (
	"encoding/json"
	"fmt"
	"library_api/entities"
	"library_api/repositories"
	"net/http"
	"sort"
	"strconv"
)

func CreateBooks(w http.ResponseWriter, r *http.Request) { // TODO discuss about create action
	name := r.FormValue("name")
	author := r.FormValue("author")
	year, err := strconv.Atoi(r.FormValue("year"))

	if err != nil {
		fmt.Println("year incorrect")
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	if name == "" || author == "" {
		fmt.Println("name or author incorrect")
		w.WriteHeader(http.StatusBadRequest)
	} else {
		books := []entities.Book{
			{name, author, year},
		}

		book_repository := repositories.NewBookRepository()

		_, err := book_repository.Create(books)
		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully Create books")
		w.WriteHeader(http.StatusCreated)
	}
}

func ListOfBooks(w http.ResponseWriter, r *http.Request) {

	initial_books := []entities.Book{
		{"Rage", "Stephen King", 1977},
		{"Philosopher's Stone", "J. K. Rowling", 1997},
		{"All Quiet on the Western Front", "Erich Maria Remarque", 1929},
	}

	book_repository := repositories.NewBookRepository()

	_, err := book_repository.Create(initial_books)
	if err != nil {
		panic(err)
	}

	books, err := book_repository.List()
	if err != nil {
		panic(err)
	}

	sort.Sort(entities.ByYear(books))

	res, err := json.Marshal(books) // TODO think how convert year
	if err != nil {
		panic(err)
	}

	// output as json file
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}
