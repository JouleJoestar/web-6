package main

// некоторые импорты нужны для проверки
import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/getMessage", handleGetRequest)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, web!")
}

/*- Create (POST)
- Read (GET)
- Update (PUT/PATCH)
- Delete (DELETE)*/
