package usecase

import (
	"github.com/olumide95/go-library/domain"
	"github.com/olumide95/go-library/models"
)

type bookUsecase struct {
	bookRepository models.BookRepository
}

func NewbookUsecase(bookRepository models.BookRepository) domain.BookUsecase {
	return &bookUsecase{
		bookRepository: bookRepository,
	}
}

func (bu *bookUsecase) AllBooks() ([]models.Book, error) {
	return bu.bookRepository.All()
}

func (bu *bookUsecase) Create(book *models.Book) error {
	return bu.bookRepository.Create(book)
}

func (bu *bookUsecase) BorrowBook(id uint, book *models.Book) error {
	return bu.bookRepository.Update(id, book)
}

func (bu *bookUsecase) ReturnBook(id uint, book *models.Book) error {
	return bu.bookRepository.Update(id, book)
}

func (bu *bookUsecase) UpdateBookQuantity(id uint, book *models.Book) error {
	return bu.bookRepository.Update(id, book)
}

func (bu *bookUsecase) CreateBulk(books *[]models.Book) error {
	return bu.bookRepository.CreateBulk(books)
}

func (bu *bookUsecase) GetBookByID(id uint) (models.Book, error) {
	return bu.bookRepository.GetByID(id)
}

func (bu *bookUsecase) Delete(id []uint) error {
	return bu.bookRepository.Delete(id)
}
