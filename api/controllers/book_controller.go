package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olumide95/go-library/api/util"
	"github.com/olumide95/go-library/domain"
	"github.com/olumide95/go-library/models"
)

type BookController struct {
	BookUsecase domain.BookUsecase
}

func (bc *BookController) AllBooks(c *gin.Context) {
	var books []models.Book
	books, err := bc.BookUsecase.AllBooks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"books": util.ErrorResponse{Message: err.Error()}})
	}

	c.JSON(http.StatusOK, util.SuccessResponse{Message: "Books retrived Successfully!", Data: books})
}

func (bc *BookController) StoreBooks(c *gin.Context) {
	var request *domain.StoreBooksRequest

	err := c.ShouldBind(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse{Message: err.Error()})
		return
	}

	var books []models.Book

	for _, val := range *request {
		books = append(books, models.Book{Title: val.Title, Author: val.Author, Quantity: val.Quantity})
	}

	err = bc.BookUsecase.CreateBulk(&books)

	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, util.SuccessResponse{Message: "Books Created Successfully!"})
}

func (bc *BookController) UpdateBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (bc *BookController) DeleteBooks(c *gin.Context) {
	var request *domain.DeleteBooksRequest

	err := c.ShouldBind(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse{Message: err.Error()})
		return
	}

	var IDs []uint
	for _, val := range *request {
		IDs = append(IDs, val.ID)
	}

	RowsAffected, err := bc.BookUsecase.Delete(IDs)

	if err != nil || RowsAffected == 0 {
		c.JSON(http.StatusNotFound, util.ErrorResponse{Message: "Error Deleting Books."})
		return
	}

	c.JSON(http.StatusOK, util.SuccessResponse{Message: "Book Deleted Successfully!"})
}

func (bc *BookController) BorrowBook(c *gin.Context) {
	var request *domain.BorrowBookRequest

	err := c.ShouldBind(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse{Message: err.Error()})
		return
	}

	bookBorrowed := bc.BookUsecase.BorrowBook(request.ID, 1)

	if !bookBorrowed {
		c.JSON(http.StatusNotFound, util.ErrorResponse{Message: "Error Borrowing Book."})
		return
	}

	c.JSON(http.StatusOK, util.SuccessResponse{Message: "Book Borrowed Successfully!", Data: bookBorrowed})
}

func (bc *BookController) ReturnBook(c *gin.Context) {
	var request *domain.ReturnBookRequest

	err := c.ShouldBind(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse{Message: err.Error()})
		return
	}

	bookBorrowed := bc.BookUsecase.ReturnBook(request.ID, request.LogID)

	if !bookBorrowed {
		c.JSON(http.StatusNotFound, util.ErrorResponse{Message: "Error Returning Book."})
		return
	}

	c.JSON(http.StatusOK, util.SuccessResponse{Message: "Book Returned Successfully!", Data: bookBorrowed})
}