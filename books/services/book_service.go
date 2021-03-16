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
	Update(user *books.Book) (*books.Book, rest_errors.RestErr)
	Delete(id int64) rest_errors.RestErr
	FindByMaxPrice(price int64) ([]*books.Book, rest_errors.RestErr)
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

func (s *BooksService) UpdateBook(isPartial bool, book *books.Book) (*books.Book, rest_errors.RestErr) {

	current, err := s.BRepository.Get(book.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if book.Name != "" {
			current.Name = book.Name
		}

		if book.Description != "" {
			current.Description = book.Description
		}

		if book.Genre != "" {
			current.Genre = book.Genre
		}

		if book.Price != 0 {
			current.Price = book.Price
		}
	} else {
		current.Name = book.Name
		current.Description = book.Description
		current.Genre = book.Genre
		current.Price = book.Price
	}
	current, err = s.BRepository.Update(current)
	if err != nil {
		return nil, err
	}
	return current, nil
}

func (s *BooksService) DeleteBook(bookId int64) rest_errors.RestErr {
	return s.BRepository.Delete(bookId)
}

func (s *BooksService) SearchBook(price int64) (books.Books, rest_errors.RestErr) {
	return s.BRepository.FindByMaxPrice(price)
}
