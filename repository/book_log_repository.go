package repository

import (
	"log"

	"github.com/olumide95/go-library/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type bookLogRepository struct {
	database *gorm.DB
}

func NewBookLogRepository(DB *gorm.DB) models.BookLogRepository {
	return &bookLogRepository{DB}
}

func (blr *bookLogRepository) Create(bookLog *models.BookLog) error {

	result := blr.database.Create(bookLog)

	return result.Error
}

func (blr *bookLogRepository) Update(id uint, bookLog *models.BookLog) (int64, error) {

	result := blr.database.Where("id = ?", id).Updates(bookLog)

	return result.RowsAffected, result.Error
}

func (blr *bookLogRepository) GetForUpdate(id uint, userId uint) (models.BookLog, error) {
	var bookLog models.BookLog
	result := blr.database.Clauses(clause.Locking{Strength: "UPDATE"}).First(&bookLog, "id = ? and user_id = ?", id, userId)

	return bookLog, result.Error
}

func (blr *bookLogRepository) GetIDsByUserId(userId uint) ([]models.BookLog, error) {
	var bookLog []models.BookLog
	result := blr.database.Where("user_id = ?", userId).Pluck("book_id", &bookLog)

	return bookLog, result.Error
}

func (blr *bookLogRepository) Delete(id []uint) error {
	var bookLog models.BookLog
	result := blr.database.Delete(&bookLog, id)

	return result.Error
}

func (blr *bookLogRepository) DeleteByBookIds(bookIds []uint) error {
	var bookLog models.BookLog
	result := blr.database.Where("book_id IN ?", bookIds).Delete(&bookLog)

	return result.Error
}

func (blr *bookLogRepository) All() ([]models.BookLog, error) {
	var bookLogs []models.BookLog
	result := blr.database.Find(&bookLogs)

	return bookLogs, result.Error
}

func (blr *bookLogRepository) WithTrx(trxHandle *gorm.DB) models.BookLogRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return blr
	}
	blr.database = trxHandle
	return blr
}
