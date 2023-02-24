package router

import (
	"library_api/controller"
	"library_api/repository"
	"net/http"

	"github.com/gorilla/mux"
)

func Route() {
	b := controller.Book{Book_repository: repository.NewBookRepository()}
	rtr := mux.NewRouter()
	rtr.HandleFunc("/create_book", b.CreateBooks).Methods(http.MethodPost)
	rtr.HandleFunc("/books", b.List).Methods(http.MethodGet)
  
	http.Handle("/", rtr)
  
	http.ListenAndServe(":8082", nil)
  }
  