package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create the route handler listening on '/'
	http.HandleFunc("/", Home)
	fmt.Println("Starting server on port 8080")

	// Start the sever
	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	msg := "This is a test for ci-cd"
	w.Write([]byte(fmt.Sprintf(msg)))
}
