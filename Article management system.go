//Article management system using RestApi
package main

import (
	"encoding/json"
	"math/rand"

	//route handlers
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Defining the model(structure)
type Article struct {
	ID      string `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

var articles []Article

//to get all the articles
func getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

//Create a new article
func createArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var article Article
	_ = json.NewDecoder(r.Body).Decode(&article)
	article.ID = strconv.Itoa(rand.Intn(1000000))
	articles = append(articles, article)
	json.NewEncoder(w).Encode(&article)
}

//to get a specific article using id
func getArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range articles {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Article{})
}

//to update an article using id
func updateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range articles {
		if item.ID == params["id"] {
			articles = append(articles[:index], articles[index+1:]...)
			var article Article
			_ = json.NewDecoder(r.Body).Decode(&article)
			article.ID = params["id"]
			articles = append(articles, article)
			json.NewEncoder(w).Encode(&article)
			return
		}
	}
	json.NewEncoder(w).Encode(articles)
}

//to delete an article using id
func deleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range articles {
		if item.ID == params["id"] {
			articles = append(articles[:index], articles[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(articles)
}
func main() {
	//Initializing the router
	router := mux.NewRouter()
	articles = append(articles, Article{ID: "01", Title: "Programming in C", Content: "C is a powerful general-purpose programming language. It can be used to develop software like operating systems, databases, compilers, and so on."})
	articles = append(articles, Article{ID: "02", Title: "Programming in C++", Content: "C++ is a powerful general-purpose programming language. It can be used to develop operating systems, browsers, games, and so on."})
	articles = append(articles, Article{ID: "03", Title: "Programming in Python", Content: "Python is a powerful general-purpose programming language. It is used in web development, data science, creating software prototypes, and so on."})
	articles = append(articles, Article{ID: "04", Title: "Programming in Golang", Content: "Go (also known as Golang) is an open source programming language developed by Google. It is a statically-typed compiled language. Go supports concurrent programming, i.e. it allows running multiple processes simultaneously."})
	//Creating endpoints
	router.HandleFunc("/articles", getArticles).Methods("GET")
	router.HandleFunc("/articles", createArticle).Methods("POST")
	router.HandleFunc("/articles/{id}", getArticle).Methods("GET")
	router.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
	router.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
