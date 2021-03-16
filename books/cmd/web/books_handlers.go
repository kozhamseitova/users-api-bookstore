package main

import (
	"github.com/gin-gonic/gin"
	books "gitlab.com/tleuzhan13/bookstore/users-api/books/domain"
	"gitlab.com/tleuzhan13/bookstore/users-api/foundation/rest_errors"
	"net/http"
	"strconv"
)

type BooksServiceInterface interface {
	GetBook(int64) (*books.Book, rest_errors.RestErr)
	AddBook(*books.Book) (*books.Book, rest_errors.RestErr)
}

func getBookId(bookIdParam string) (int64, rest_errors.RestErr) {
	bookId, bookErr := strconv.ParseInt(bookIdParam, 10, 64)
	if bookErr != nil {
		return 0, rest_errors.NewBadRequestError("user id should be a number")
	}
	return bookId, nil
}

func (a *app) Add(c *gin.Context) {

	book := &books.Book{}
	if err := c.ShouldBindJSON(book); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, saveErr := a.service.AddBook(book)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall())
}

func (a *app) Get(c *gin.Context) {
	bookId, idErr := getBookId(c.Param("book_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	var book books.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, err := a.service.GetBook(bookId)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, result.Marshall())
}
