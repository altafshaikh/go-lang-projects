package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Post struct {
	ID     string `json:id`
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

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(dbPosts)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var newPost Post
	_ = json.NewDecoder(r.Body).Decode(&newPost)

	newPost.ID = strconv.Itoa(rand.Intn(100000000))
	dbPosts = append(dbPosts, newPost)

	json.NewEncoder(w).Encode(dbPosts)
}

func getPostByUid(w http.ResponseWriter, r *http.Request) {}

func updatePostByUid(w http.ResponseWriter, r *http.Request) {}

func deletePostByUid(w http.ResponseWriter, r *http.Request) {}
