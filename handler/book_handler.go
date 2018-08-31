package handler

import (
	"encoding/json"
	"net/http"

	"github.com/viniciusou/cloud-native-go/model"
	"github.com/viniciusou/cloud-native-go/repository"
)

//BooksHandleFunc handles requests for Book API
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		writeJSON(w, books)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method"))
	}
}

//AllBooks return all books from repository
func AllBooks() []model.Book {
	values := make([]model.Book, len(repository.Books))
	idx := 0
	for _, book := range repository.Books {
		values[idx] = book
		idx++
	}

	return values
}

func writeJSON(w http.ResponseWriter, i interface{}) {
	byteSlice, err := json.Marshal(i)
	if err != nil {
		http.Error(w, "Error to serialize JSON "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json: charset=utf-8")
	w.Write(byteSlice)
}
