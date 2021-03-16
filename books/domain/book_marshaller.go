package domain

import "encoding/json"

type PublicBook struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	Price       int64  `json:"price"`
}

func (books Books) Marshall() []interface{} {
	result := make([]interface{}, len(books))
	for index, book := range books {
		result[index] = book.Marshall()
	}
	return result
}

func (book *Book) Marshall() interface{} {

	bookJson, _ := json.Marshal(book)
	var publicBook PublicBook
	json.Unmarshal(bookJson, &publicBook)
	return publicBook
}
