package services

import (
	books "gitlab.com/tleuzhan13/bookstore/users-api/books/domain"
	"gitlab.com/tleuzhan13/bookstore/users-api/foundation/rest_errors"
)

type BooksService struct {
	BRepository BRepository
}

type BRepository interface {
	Get(id int64) (*books.Book, rest_errors.RestErr)
	Save(user *books.Book) (*books.Book, rest_errors.RestErr)
}

func (s *BooksService) GetBook(bookId int64) (*books.Book, rest_errors.RestErr) {
	book, err := s.BRepository.Get(bookId)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *BooksService) AddBook(book *books.Book) (*books.Book, rest_errors.RestErr) {
	if err := book.Validate(); err != nil {
		return nil, err
	}

	userSaved, err := s.BRepository.Save(book)
	if err != nil {
		return nil, err
	}
	return userSaved, nil
}
