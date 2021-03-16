package domain

import (
	"gitlab.com/tleuzhan13/bookstore/users-api/foundation/rest_errors"
	"strings"
)

type Book struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	Price       int64  `json:"price"`
}

type Books []*Book

func (book *Book) Validate() rest_errors.RestErr {
	book.Name = strings.TrimSpace(book.Name)
	if book.Name == "" {
		return rest_errors.NewBadRequestError("invalid email address")
	}

	book.Description = strings.TrimSpace(book.Description)

	book.Genre = strings.TrimSpace(book.Genre)

	return nil
}
