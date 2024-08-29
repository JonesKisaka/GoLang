package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mux-main"
	"net/http"
)

// struct that contains info on the article type that will be accessed from the Server
type Article struct {
	Title       string "Title"
	Description string "Desc"
	Content     string "Content"
}

type Writer struct {
	Name string "Name"
	Age  string "Age"
	Year string "Year"
}

type Actor struct {
	Name  string "name"
	Age   string "age"
	Movie string "movie"
}

type Genre struct {
	Name     string "Genre"
	prodYear string "prodYear"
}

// slice of articles called Article
type Articles []Article
type Writers []Writer
type Actors []Actor
type Genres []Genre

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Rational Male", Description: "By Rollo Tomassi", Content: "The Players Handbook"},
		Article{Title: "Holy Bible", Description: "Gideons International", Content: "The Good News Bible"},
		Article{Title: "Mastery", Description: "By Robert Greene", Content: "Journey to Mastery to any thing in Life"},
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint!")
	fmt.Fprintf(w, "The Home Endpoint!\n\n\n")
	json.NewEncoder(w).Encode(articles)
}

func allWriters(w http.ResponseWriter, r *http.Request) {
	writers := Writers{
		Writer{Name: "Rollo Tomassi", Age: "45", Year: "1998"},
		Writer{Name: "Myron Gaines", Age: "35", Year: "2010"},
		Writer{Name: "Robert Greene", Age: "50", Year: "1960"},
	}
	fmt.Println("Endpoint Hit: All Writers Endpoint!")
	fmt.Fprintf(w, "The Writers Endpoint!\n\n\n")
	json.NewEncoder(w).Encode(writers)
}

func allActors(w http.ResponseWriter, r *http.Request) {
	actors := Actors{
		Actor{Name: "Jones", Age: "25", Movie: "Fast and The Furious"},
		Actor{Name: "Noel", Age: "24", Movie: "Despicable Me"},
		Actor{Name: "Monica", Age: "31", Movie: "Carolina"},
	}
	fmt.Println("Endpoint Hit: All Actors Endpoint")
	fmt.Fprintf(w, "All Actors are below!\n\n\n")
	json.NewEncoder(w).Encode(actors)
}

func allGenres(w http.ResponseWriter, r *http.Request) {
	genre := Genres{
		Genre{Name: "Action", prodYear: "1994"},
		Genre{Name: "Comedy", prodYear: "2023"},
		Genre{Name: "Horror", prodYear: "2004"},
	}

	fmt.Println("All Genres Endpoint HIt!")
	fmt.Fprintf(w, "All Genres Endpoint Hit!\n\n\n")
	json.NewEncoder(w).Encode(genre)
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POST Endpoint Hit. \n\n")
}

// the hompepage function handles all the requests to the root URL.(Uniform Resorce Locator)
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Endpoint Hit. \nNa pia tuko site mzee!\n")

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	myRouter.HandleFunc("/writers", allWriters)
	myRouter.HandleFunc("/actors", allActors)
	myRouter.HandleFunc("/genres", allGenres)
	fmt.Println("Starting server at port 8081")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}
