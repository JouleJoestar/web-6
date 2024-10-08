package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/user", handleUserRequest)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleUserRequest(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Name parameter is required", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello, %s!", name)
}
