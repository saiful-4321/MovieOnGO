package main

import "github.com/gorilla/mux"

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Name     string `json:"name"`
	Position string `json:"position"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", GetMovie).Methods("GET")
	r.HandleFunc("/movie", CreateMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", UpdateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", DeleteMovie).Methods("DELETE")
}
