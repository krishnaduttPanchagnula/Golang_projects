package main

import (
	"encoding/json" // To encode the data into json when sending a response
	"fmt"           // to print out and scan things
	"log"           // to create logs for api
	"math/rand"     // to create ID to the input data
	"net/http"      // to create a web server

	"strconv" // to convert the data into strings

	"github.com/gorilla/mux" //Creates request router
)

type Movie struct {
	ID          string `json:id`
	Genre       string `json:Genre`
	Budget      int    `json:budget`
	Title       string `json:title`
	Technicians `json:Technicians`
}

type Technicians struct {
	Hero    string `json:hero`
	Heroine string `json:Heroine`
}

var movies []Movie

func getmovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deletemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break

		}
		json.NewEncoder(w).Encode(movies)
		return
	}
}

func getmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range movies {

		if item.ID == params["id"] {

			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func createMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newMovie Movie

	_ = json.NewDecoder(r.Body).Decode(&newMovie)

	newMovie.ID = strconv.Itoa(rand.Intn(1000000))

	movies = append(movies, newMovie)
	json.NewEncoder(w).Encode(movies)

}

func updateMovies(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	for index, item := range movies {

		if item.ID == params["id"] {

			movies = append(movies[:index], movies[index+1:]...)

			var updatedMovie Movie
			_ = json.NewDecoder(r.Body).Decode(&updatedMovie)

			updatedMovie.ID = strconv.Itoa(rand.Intn(1000000))

			movies = append(movies, updatedMovie)

			json.NewEncoder(w).Encode(movies)

		}
	}

}

func main() {

	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Genre: "horror", Budget: 500000, Title: "SAW", Technicians: Technicians{Hero: "JONNY DEPP", Heroine: "Scarlet"}})

	movies = append(movies, Movie{ID: "2", Genre: "Romance", Budget: 1000000, Title: "TEST", Technicians: Technicians{Hero: "BALE", Heroine: "ARMAS"}})

	fmt.Println(movies)

	r.HandleFunc("/movies", getmovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getmovie).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movie/{id}", deletemovie).Methods("DELETE")
	fmt.Println("Starting the Server...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
