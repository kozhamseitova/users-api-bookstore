package main

import (
	"github.com/gin-gonic/gin"
	books "gitlab.com/tleuzhan13/bookstore/users-api/books/domain"
	"gitlab.com/tleuzhan13/bookstore/users-api/foundation/rest_errors"
	users "gitlab.com/tleuzhan13/bookstore/users-api/users/domain"
	"net/http"
	"strconv"
)

type BooksServiceInterface interface {
	GetBook(int64) (*books.Book, rest_errors.RestErr)
	AddBook(*books.Book) (*books.Book, rest_errors.RestErr)
	UpdateBook(bool, *books.Book) (*books.Book, rest_errors.RestErr)
	DeleteBook(int64) rest_errors.RestErr
	SearchBook(int64) (books.Book, rest_errors.RestErr)
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
	userId, idErr := getBookId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, err := a.service.GetBook(userId)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, result.Marshall())
}
