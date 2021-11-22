package main

import (
	"math/rand"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
)

// declare struct movie and director

type Movie struct {
	Id string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director    `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName string  `json:"lastname"`
}

var movies []Movie // declaring a slice that will contain all movies

func getMovies(w http.ResponseWriter , r*http.Request ){
	w.Header().Set("content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(movies)
	
}

func getMovie(w http.ResponseWriter , r * http.Request) {
	w.Header().Set("content-Type" , "application/json")
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)

	for _ , item := range movies {
		if item.Id == params["id"]{
			encoder.Encode(item)
			return 
		}
	}
}

func deleteMovie(w http.ResponseWriter , r *http.Request) {
	w.Header().Set("content-Type" , "application/json")
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)

	for index , item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index] , movies[index+1:]...)
			break
		}
	}
	encoder.Encode(movies)

}


func addMovie(w http.ResponseWriter , r *http.Request) {
	w.Header().Set("content-Type" , "application/json")
	var newMovie Movie 
	decoder := json.NewDecoder(r.Body) 
	decoder.Decode(&newMovie)

	newMovie.Id = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies , newMovie)
	encoder := json.NewEncoder(w)
	encoder.Encode(movies) 
}

func updateMovie(w http.ResponseWriter , r *http.Request) {
	w.Header().Set("content-Type" , "application/json") 
	params := mux.Vars(r) 
	encoder := json.NewEncoder(w) 


	for index , item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index] , movies[index+1:]...)
			var updatedMovie Movie
			json.NewDecoder(r.Body).Decode(&updatedMovie)
			updatedMovie.Id = params["id"]

			movies = append(movies , updatedMovie) 	
		}
	}

	encoder.Encode(movies)

}





func main() {
	firstMovie := Movie{Id : "1" , Isbn: "b45" , Title :"The story of my Life" , Director: &Director{FirstName: "Dev" , LastName: "Prakash"}}
	secondMovie := Movie{Id : "2" , Isbn : "c64", Title: "Logan" , Director: &Director{FirstName : "James" , LastName: "Mangold" }}
	movies = append(movies , firstMovie)
	movies = append(movies , secondMovie)


	r := mux.NewRouter() // creates an instance of the router
	r.HandleFunc("/movies" ,getMovies ).Methods("GET")
	r.HandleFunc("/movies/{id}" , getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}" , deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies", addMovie).Methods("POST") 
	r.HandleFunc("/movies/{id}" , updateMovie).Methods("PUT")
	fmt.Printf("Server starting at port 4000\n")
	err := http.ListenAndServe(":4000" , r )
	if err != nil {
		log.Fatal(err)
	}

}