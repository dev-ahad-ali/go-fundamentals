package main

import (
	"encoding/json"
	"net/http"
)

// Post defines the structure of our resource
type Post struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func getPostHandler(w http.ResponseWriter, r *http.Request) {
	// Create a mock resource state
	post := Post{ID: 42, Title: "Into The Backend", Author: "Habib"}

	// Set the content type to json
	w.Header().Set("Content-Type", "application/json")

	// Encode and transfer the resource state to the client
	json.NewEncoder(w).Encode(post)
}

func main() {
	http.HandleFunc("/post", getPostHandler)
	http.ListenAndServe(":8080", nil)
}
