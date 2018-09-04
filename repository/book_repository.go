package repository

import (
	"encoding/json"
	"fmt"

	"github.com/viniciusou/cloud-native-go/model"
)

//Books mock slice simulating existing data from a database
var Books = map[string]model.Book{
	"9876543210": model.Book{Title: "Go microservices", Author: "John", ISBN: "9876543210"},
	"0123456789": model.Book{Title: "Go cloud native", Author: "Mary", ISBN: "0123456789"},
}

//ReadJSON deserializes JSON data to Book struct type
func ReadJSON(data []byte) model.Book {
	book := model.Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		fmt.Println("Error to deserialize JSON", err.Error())
	}

	return book
}

//WriteJSON serializes data to JSON
func WriteJSON(i interface{}) ([]byte, error) {
	byteSlice, err := json.Marshal(i)
	if err != nil {
		fmt.Println("Error to serialize JSON ", err.Error())
		return nil, err
	}

	return byteSlice, nil
}
