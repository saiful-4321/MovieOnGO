package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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

	// initiate seeder data
	insertMovieSeeder()

	r.HandleFunc("/movies", GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", GetMovie).Methods("GET")
	r.HandleFunc("/movie", CreateMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", UpdateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", DeleteMovie).Methods("DELETE")

	fmt.Printf("starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func insertMovieSeeder() {
	movies = append(movies, Movie{ID: "1", Isbn: "4321", Title: "dark-night", Director: &Director{Name: "Saiful Islam", Position: "Head dir"}})
	movies = append(movies, Movie{ID: "2", Isbn: "1234", Title: "purple-night", Director: &Director{Name: "Saiful Islam", Position: "Head dir"}})
	movies = append(movies, Movie{ID: "3", Isbn: "1122", Title: "Geen-night", Director: &Director{Name: "Samim Islam", Position: "Sub dir"}})
	movies = append(movies, Movie{ID: "4", Isbn: "5544", Title: "Purple-night", Director: &Director{Name: "Nusaifa Islam", Position: "OOP dir"}})
	movies = append(movies, Movie{ID: "5", Isbn: "5566", Title: "totle-night", Director: &Director{Name: "Kona Islam", Position: "DDR dir"}})
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}
