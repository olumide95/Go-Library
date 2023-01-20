package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olumide95/go-library/api/util"
	"github.com/olumide95/go-library/domain"
	"github.com/olumide95/go-library/models"
	"gorm.io/gorm"
)

type BookController struct {
	BookUsecase domain.BookUsecase
}

func (bc *BookController) AllBooks(c *gin.Context) {
	var books []models.Book
	books, _ = bc.BookUsecase.AllBooks()

	c.JSON(http.StatusOK, util.SuccessResponse{Message: "Books retrived Successfully!", Data: books})
}

func (bc *BookController) AllBorrowedBooks(c *gin.Context) {
	userId, exists := c.Get("currentUserId")

	if !exists {
		c.JSON(http.StatusNotFound, util.ErrorResponse{Message: "User not Found."})
		return
	}

	var books []models.BookLog
	books, _ = bc.BookUsecase.AllBorrowedBooks(userId.(uint))

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
	var request *domain.UpdateBookRequest

	err := c.ShouldBind(&request)
	log.Println(request.BookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse{Message: err.Error()})
		return
	}

	book := models.Book{ID: request.BookID, Title: request.Title, Author: request.Author, Quantity: request.Quantity}

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	bookUpdated := bc.BookUsecase.WithTrx(txHandle).UpdateBook(&book)

	if !bookUpdated {
		c.JSON(http.StatusInternalServerError, util.SuccessResponse{Message: "Error Updating Book!"})
		return
	}

	c.JSON(http.StatusOK, util.SuccessResponse{Message: "Books Updated Successfully!"})
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

	txHandle := c.MustGet("db_trx").(*gorm.DB)
	booksDeleted := bc.BookUsecase.WithTrx(txHandle).Delete(IDs)

	if !booksDeleted {
		c.JSON(http.StatusInternalServerError, util.ErrorResponse{Message: "Error Deleting Books."})
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

	userId, exists := c.Get("currentUserId")

	if !exists {
		c.JSON(http.StatusNotFound, util.ErrorResponse{Message: "User not Found."})
		return
	}

	txHandle := c.MustGet("db_trx").(*gorm.DB)
	bookBorrowed := bc.BookUsecase.WithTrx(txHandle).BorrowBook(request.BookID, userId.(uint))

	if !bookBorrowed {
		c.JSON(http.StatusInternalServerError, util.ErrorResponse{Message: "Error Borrowing Book."})
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

	userId, exists := c.Get("currentUserId")

	if !exists {
		c.JSON(http.StatusNotFound, util.ErrorResponse{Message: "User not Found."})
		return
	}

	txHandle := c.MustGet("db_trx").(*gorm.DB)
	bookReturned := bc.BookUsecase.WithTrx(txHandle).ReturnBook(request.LogID, userId.(uint))

	if !bookReturned {
		c.JSON(http.StatusInternalServerError, util.ErrorResponse{Message: "Error Returning Book."})
		return
	}

	c.JSON(http.StatusOK, util.SuccessResponse{Message: "Book Returned Successfully!", Data: bookReturned})
}
