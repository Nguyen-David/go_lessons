package router

import (
    "github.com/gorilla/mux"
    "net/http"
	"library_api/controllers"
)

func Route() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/create_book", controllers.CreateBooks).Methods("POST")
	rtr.HandleFunc("/books", controllers.ListOfBooks).Methods("GET")
  
	http.Handle("/", rtr)
  
	http.ListenAndServe(":8082", nil)
  }
  