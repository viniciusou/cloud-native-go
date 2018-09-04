package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/viniciusou/cloud-native-go/model"
	"github.com/viniciusou/cloud-native-go/repository"
)

//BooksHandleFunc handles requests for Book API from route "/api/books"
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		byteSlice, err := repository.WriteJSON(books)
		if err != nil {
			http.Error(w, "Error to read JSON "+err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Add("Content-Type", "application/json: charset=utf-8")
		w.Write(byteSlice)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := repository.ReadJSON(body)
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

//BookHandleFunc handles requests for Book API from route "/api/books/"
func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Path[len("/api/books/"):]

	switch method := r.Method; method {
	case http.MethodGet:
		book, found := GetBook(isbn)
		if found {
			byteSlice, err := repository.WriteJSON(book)
			if err != nil {
				http.Error(w, "Error to read JSON "+err.Error(), http.StatusBadRequest)
				return
			}
			w.Header().Add("Content-Type", "application/json: charset=utf-8")
			w.Write(byteSlice)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := repository.ReadJSON(body)
		exists := UpdateBook(isbn, book)
		if exists {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
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

//GetBook returns the book for a given ISBN
func GetBook(isbn string) (model.Book, bool) {
	book, found := repository.Books[isbn]

	return book, found
}

//UpdateBook updates an existing book
func UpdateBook(isbn string, book model.Book) bool {
	_, exists := repository.Books[isbn]
	if exists {
		repository.Books[isbn] = book
	}

	return exists
}
