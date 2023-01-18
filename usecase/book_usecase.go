package usecase

import (
	"time"

	"github.com/olumide95/go-library/domain"
	"github.com/olumide95/go-library/models"
)

type bookUsecase struct {
	bookRepository    models.BookRepository
	bookLogRepository models.BookLogRepository
}

func NewbookUsecase(bookRepository models.BookRepository, bookLogRepository models.BookLogRepository) domain.BookUsecase {
	return &bookUsecase{
		bookRepository:    bookRepository,
		bookLogRepository: bookLogRepository,
	}
}

func (bu *bookUsecase) AllBooks() ([]models.Book, error) {
	return bu.bookRepository.All()
}

func (bu *bookUsecase) Create(book *models.Book) error {
	return bu.bookRepository.Create(book)
}

func (bu *bookUsecase) BorrowBook(id uint, userId uint) bool {
	book, err := bu.bookRepository.GetByIDForUpdate(id)

	if err != nil {
		return false
	}

	if book.Quantity == 0 {
		return false
	}

	book.Quantity -= 1
	result, err := bu.bookRepository.Update(id, &book)

	if err != nil || result == 0 {
		return false
	}

	bookLog := models.BookLog{BookId: id, UserId: userId}
	err = bu.bookLogRepository.Create(&bookLog)

	if err != nil {
		return false
	}

	return true
}

func (bu *bookUsecase) ReturnBook(logId uint, userId uint) bool {
	bookLog, err := bu.bookLogRepository.GetForUpdate(logId, userId)

	if err != nil {
		return false
	}

	if bookLog.ReturnedAt != nil {
		return false
	}

	book, err := bu.bookRepository.GetByIDForUpdate(bookLog.BookId)

	if err != nil {
		return false
	}

	book.Quantity += 1
	result, err := bu.bookRepository.Update(book.ID, &book)

	if err != nil || result == 0 {
		return false
	}

	now := time.Now()
	bookLog.ReturnedAt = &now
	result, err = bu.bookLogRepository.Update(logId, &bookLog)

	if err != nil || result == 0 {
		return false
	}

	return true
}

func (bu *bookUsecase) UpdateBookQuantity(id uint, quantity uint16) (int64, error) {
	book := models.Book{Quantity: quantity}
	return bu.bookRepository.Update(id, &book)
}

func (bu *bookUsecase) CreateBulk(books *[]models.Book) error {
	return bu.bookRepository.CreateBulk(books)
}

func (bu *bookUsecase) GetBookByID(id uint) (models.Book, error) {
	return bu.bookRepository.GetByID(id)
}

func (bu *bookUsecase) Delete(id []uint) (int64, error) {
	return bu.bookRepository.Delete(id)
}
