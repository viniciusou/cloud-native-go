package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/viniciusou/cloud-native-go/handler"
)

func main() {
	http.HandleFunc("/api/echo", echo)
	http.HandleFunc("/api/books", handler.BooksHandleFunc)

	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	return ":" + port

}

func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]

	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}
