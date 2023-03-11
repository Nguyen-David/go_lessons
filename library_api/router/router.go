package router

import (
	"library_api/config"
	"library_api/controller"
	"library_api/repository"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Route(fileConnection *os.File) {
	b := controller.Book{Book_repository: repository.NewCsvRepository(fileConnection)}
	rtr := mux.NewRouter()
	rtr.HandleFunc("/create_book", b.CreateBook).Methods(http.MethodPost)
	rtr.HandleFunc("/books", b.List).Methods(http.MethodGet)
  
	http.Handle(config.ConfigGlobal.ListeningURL, rtr)
  
	http.ListenAndServe(config.ConfigGlobal.Port, nil)
  }
  