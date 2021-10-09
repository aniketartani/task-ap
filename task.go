package main

import (
	"encoding/json"
	"mux-master"
	"net/http"
	"strconv"
)

// User is a struct that represents a user in our application
type User struct {
	FullName string `json:"fullName"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Post is a struct that represents a single post
type Post struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Users struct {
	fname string `"json:"fname"`
	lname string `"json:"lname"`
}

var posts []Post = []Post{}
var user []Users = []Users{}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/posts", addItem).Methods("POST")

	router.HandleFunc("/posts", getAllPosts).Methods("GET")

	router.HandleFunc("/posts/{id}", getPost).Methods("GET")

	http.ListenAndServe(":3000", router)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	// get the ID of the post from the route parameter
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// there was an error
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	// error checking
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}

	post := posts[id]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	// get Item value from the JSON body
	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)

	posts = append(posts, newPost)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(posts)
}
func adduser(w http.ResponseWriter, r *http.Request) {
	// get Item value from the JSON body
	var newUser Users
	json.NewDecoder(r.Body).Decode(&newUser)

	posts = append(Users, newUser)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(posts)
}
