package router

import (
    "github.com/gorilla/mux"
    "net/http"
	"library_api/controller"
)

func Route() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/create_book", controller.CreateBooks).Methods(http.MethodPost)
	rtr.HandleFunc("/books", controller.List).Methods(http.MethodGet)
  
	http.Handle("/", rtr)
  
	http.ListenAndServe(":8082", nil)
  }
  