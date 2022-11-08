package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Post struct {
	Title  string `json:title`
	Body   string `json:body`
	Author User   `json:user`
}

type User struct {
	FirstName string `json:firstName`
	LastName  string `json:lastName`
	Email     string `json:email`
}

var dbPosts []Post = []Post{}

func main() {
	PORT := ":8080"
	router := mux.NewRouter()

	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", createPost).Methods("POST")
	router.HandleFunc("/posts/{id}", getPostByUid).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePostByUid).Methods("PUT")
	router.HandleFunc("/posts/{id}", deletePostByUid).Methods("DELETE")

	fmt.Println("starting server at port ", PORT)

	http.ListenAndServe(PORT, router)
}

// route handlers

func getPosts(w http.ResponseWriter, r *http.Request) {}

func createPost(w http.ResponseWriter, r *http.Request) {}

func getPostByUid(w http.ResponseWriter, r *http.Request) {}

func updatePostByUid(w http.ResponseWriter, r *http.Request) {}

func deletePostByUid(w http.ResponseWriter, r *http.Request) {}
