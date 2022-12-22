package router

import (
	"github.com/gorilla/mux"
	"github.com/saurabhkanawade/controller"
)

func Router() *mux.Router {

	route := mux.NewRouter()

	route.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	route.HandleFunc("/api/movies", controller.CreateMovie).Methods("POST")
	route.HandleFunc("/api/movies/{id}", controller.DeleteOneMovie).Methods("DELETE")
	route.HandleFunc("/api/movies", controller.DeleteAllMovie).Methods("DELETE")
	route.HandleFunc("/api/movies/{id}", controller.MarkAdWatched).Methods("PUT")

	return route
}
