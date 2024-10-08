package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var counter int = 0

func main() {
	http.HandleFunc("/count", handleCount)
	fmt.Println("Сервер запущен на порту :3333")
	http.ListenAndServe(":3333", nil)
}

func handleCount(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "%d", counter)
	case "POST":
		count, err := strconv.Atoi(r.FormValue("count"))
		if err != nil {
			http.Error(w, "это не число", http.StatusBadRequest)
			return
		}
		counter += count
		fmt.Fprintf(w, "Success")
	default:
		http.Error(w, "Неизвестный метод", http.StatusMethodNotAllowed)
	}
}
