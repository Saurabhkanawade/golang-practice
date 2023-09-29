package router

import (
	"github.com/SaurabhKanawade/BookAuthor/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	route := mux.NewRouter()

	route.HandleFunc("/api/books", controller.CreateBook).Methods("POST")

	return route
}
