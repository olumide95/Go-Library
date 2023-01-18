package domain

import (
	"github.com/olumide95/go-library/models"
)

type BookUsecase interface {
	Create(book *models.Book) error
	UpdateBookQuantity(id uint, book *models.Book) error
	ReturnBook(id uint, book *models.Book) error
	CreateBulk(books *[]models.Book) error
	GetBookByID(id uint) (models.Book, error)
	Delete(id []uint) error
	AllBooks() ([]models.Book, error)
}
