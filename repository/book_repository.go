package repository

import (
	"github.com/olumide95/go-library/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type bookRepository struct {
	database *gorm.DB
}

func NewBookRepository(DB *gorm.DB) models.BookRepository {
	return &bookRepository{DB}
}

func (br *bookRepository) Create(book *models.Book) error {

	result := br.database.Create(book)

	return result.Error
}

func (br *bookRepository) CreateBulk(book *[]models.Book) error {

	result := br.database.CreateInBatches(book, 100)

	return result.Error
}

func (br *bookRepository) Update(id uint, book *models.Book) (int64, error) {

	result := br.database.Where("id = ?", id).Updates(book)

	return result.RowsAffected, result.Error
}

func (br *bookRepository) GetByIds(ids []uint) ([]models.Book, error) {
	var book []models.Book
	result := br.database.First(&book, "id IN ?", ids)

	return book, result.Error
}

func (br *bookRepository) GetByID(id uint) (models.Book, error) {
	var book models.Book
	result := br.database.First(&book, "id = ?", id)

	return book, result.Error
}

func (br *bookRepository) GetByIDForUpdate(id uint) (models.Book, error) {
	var book models.Book
	result := br.database.Clauses(clause.Locking{Strength: "UPDATE"}).First(&book, "id = ?", id)

	return book, result.Error
}

func (br *bookRepository) Delete(id []uint) (int64, error) {
	var book models.Book
	result := br.database.Delete(&book, id)

	return result.RowsAffected, result.Error
}

func (br *bookRepository) All() ([]models.Book, error) {
	var books []models.Book
	result := br.database.Find(&books)

	return books, result.Error
}
