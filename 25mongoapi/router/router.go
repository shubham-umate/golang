package router

import (
	"github.com/gorilla/mux"
	"github.com/shubham-umate/mongoapi/controller"
)

func Router() *mux.Router {
	r := mux.NewRouter()
		r.HandleFunc("/api/movies",controller.GetMyAllMovies).Methods("GET")
		r.HandleFunc("/api/movie",controller.CreateMovie).Methods("POST")
		r.HandleFunc("/api/movie/{id}",controller.MarkAsWatched).Methods("PUT")
		r.HandleFunc("/api/movie/{id}",controller.DeleteOneMovie).Methods("DELETE")
		r.HandleFunc("/api/movies",controller.DeleteAllMovies).Methods("DELETE")
		



	return r
}