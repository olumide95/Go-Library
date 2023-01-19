package domain

import (
	"github.com/olumide95/go-library/models"
)

type StoreBooksRequest []struct {
	Title    string `form:"title" binding:"required"`
	Author   string `form:"author" binding:"required"`
	Quantity uint16 `form:"quantity" binding:"required"`
}

type UpdateBookRequest struct {
	ID       uint   `form:"bookId" binding:"required"`
	Title    string `form:"title"`
	Author   string `form:"author"`
	Quantity uint16 `form:"quantity"`
}

type BorrowBookRequest struct {
	BookID uint `form:"bookId" binding:"required"`
}

type ReturnBookRequest struct {
	LogID uint `form:"logId" binding:"required"`
}

type DeleteBooksRequest []struct {
	ID uint `form:"id" binding:"required"`
}

type BookUsecase interface {
	Create(book *models.Book) error
	UpdateBook(book *models.Book) bool
	BorrowBook(id uint, userId uint) bool
	ReturnBook(logId uint, userId uint) bool
	CreateBulk(books *[]models.Book) error
	GetBookByID(id uint) (models.Book, error)
	Delete(id []uint) bool
	AllBooks() ([]models.Book, error)
	AllBorrowedBooks(userId uint) ([]models.Book, error)
}
