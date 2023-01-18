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

func (ur *bookRepository) Create(book *models.Book) error {

	result := ur.database.Create(book)

	return result.Error
}

func (ur *bookRepository) CreateBulk(book *[]models.Book) error {

	result := ur.database.CreateInBatches(book, 100)

	return result.Error
}

func (ur *bookRepository) Update(id uint, book *models.Book) (int64, error) {

	result := ur.database.Where("id = ?", id).Updates(book)

	return result.RowsAffected, result.Error
}

func (ur *bookRepository) GetByID(id uint) (models.Book, error) {
	var book models.Book
	result := ur.database.First(&book, "id = ?", id)

	return book, result.Error
}

func (ur *bookRepository) GetByIDForUpdate(id uint) (models.Book, error) {
	var book models.Book
	result := ur.database.Clauses(clause.Locking{Strength: "UPDATE"}).First(&book, "id = ?", id)

	return book, result.Error
}

func (ur *bookRepository) Delete(id []uint) (int64, error) {
	var book models.Book
	result := ur.database.Delete(&book, id)

	return result.RowsAffected, result.Error
}

func (ur *bookRepository) All() ([]models.Book, error) {
	var books []models.Book
	result := ur.database.Find(&books)

	return books, result.Error
}
