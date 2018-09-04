package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/viniciusou/cloud-native-go/model"
)

func TestWriteJSON(t *testing.T) {
	book := model.Book{Title: "Go microservices", Author: "John", ISBN: "9876543210"}
	json, _ := WriteJSON(book)

	assert.Equal(t, `{"title":"Go microservices","author":"John","isbn":"9876543210"}`, string(json), "Write JSON marshalling wrong")
}

func TestReadJSON(t *testing.T) {
	json := []byte(`{"title":"Go microservices","author":"John","isbn":"9876543210"}`)
	book := ReadJSON(json)

	assert.Equal(t, model.Book{Title: "Go microservices", Author: "John", ISBN: "9876543210"}, book, "Read JSON unmarshalling wrong")
}
