package domain

import (
	"github.com/olumide95/go-library/models"
)

type StoreBooksRequest []struct {
	Title    string `form:"title" binding:"required"`
	Author   string `form:"author" binding:"required"`
	Quantity uint16 `form:"quantity" binding:"required"`
}

type BorrowBookRequest struct {
	ID uint `form:"id" binding:"required"`
}

type ReturnBookRequest struct {
	ID    uint `form:"id" binding:"required"`
	LogID uint `form:"log_id" binding:"required"`
}

type DeleteBooksRequest []struct {
	ID uint `form:"id" binding:"required"`
}

type BookUsecase interface {
	Create(book *models.Book) error
	UpdateBookQuantity(id uint, quantity uint16) (int64, error)
	BorrowBook(id uint, userId uint) bool
	ReturnBook(id uint, logId uint) (int64, error)
	CreateBulk(books *[]models.Book) error
	GetBookByID(id uint) (models.Book, error)
	Delete(id []uint) (int64, error)
	AllBooks() ([]models.Book, error)
}
