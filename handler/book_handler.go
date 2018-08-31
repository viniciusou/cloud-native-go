package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := readJSON(body)
		isbn, created := CreateBook(book)
		if created {
			w.Header().Add("location", "/api/books/"+isbn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
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

//CreateBook creates a new book if it does not exist
func CreateBook(book model.Book) (string, bool) {
	_, exists := repository.Books[book.ISBN]
	if exists {
		return "", false
	}

	repository.Books[book.ISBN] = book
	return book.ISBN, true
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

func readJSON(data []byte) model.Book {
	book := model.Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		fmt.Println("Error to deserialize JSON", err.Error())
	}

	return book
}
