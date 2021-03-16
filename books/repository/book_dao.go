package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	books "gitlab.com/tleuzhan13/bookstore/users-api/books/domain"
	"gitlab.com/tleuzhan13/bookstore/users-api/foundation/rest_errors"
	"log"
)

const (
	queryInsertBook = "INSERT INTO books(name, description, genre, price) VALUES ($1,$2,$3,$4) RETURNING id"
	queryGetBook    = "SELECT id, name, description, genre, price FROM books WHERE id=$1"
)

type BookRepository struct {
	Pool     *pgxpool.Pool
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

func (r *books.Book) Get(id int64) (*books.Book, rest_errors.RestErr) {
	row := r.Pool.QueryRow(context.Background(), queryGetBook, id)

	var book *books.Book
	book = new(books.Book)

	err := row.Scan(&book.Id, &book.Name, &book.Description, &book.Genre, &book.Price)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("internal server error", err)
	}
	return book, nil
}

func (r *BookRepository) Save(book *books.Book) (*books.Book, rest_errors.RestErr) {
	var bookId int64
	row := r.Pool.QueryRow(context.Background(), queryInsertBook, book.Name, book.Description, book.Genre,
		book.Price)
	err := row.Scan(&bookId)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("internal server error", err)
	}
	book.Id = bookId
	return book, nil
}
