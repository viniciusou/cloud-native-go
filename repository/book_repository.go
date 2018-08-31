package repository

import "github.com/viniciusou/cloud-native-go/model"

//Books mock slice simulating existing data from a database
var Books = map[string]model.Book{
	"9876543210": model.Book{Title: "Go microservices", Author: "John", ISBN: "9876543210"},
	"0123456789": model.Book{Title: "Go cloud native", Author: "Mary", ISBN: "0123456789"},
}
