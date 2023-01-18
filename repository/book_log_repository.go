package repository

import (
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

func (ur *bookLogRepository) Create(bookLog *models.BookLog) error {

	result := ur.database.Create(bookLog)

	return result.Error
}

func (ur *bookLogRepository) Update(id uint, bookLog *models.BookLog) (int64, error) {

	result := ur.database.Where("id = ?", id).Updates(bookLog)

	return result.RowsAffected, result.Error
}

func (ur *bookLogRepository) GetForUpdate(id uint, userId uint) (models.BookLog, error) {
	var bookLog models.BookLog
	result := ur.database.Clauses(clause.Locking{Strength: "UPDATE"}).First(&bookLog, "id = ? and user_id = ?", id, userId)

	return bookLog, result.Error
}

func (ur *bookLogRepository) Delete(id []uint) error {
	var bookLog models.BookLog
	result := ur.database.Delete(&bookLog, id)

	return result.Error
}

func (ur *bookLogRepository) All() ([]models.BookLog, error) {
	var bookLogs []models.BookLog
	result := ur.database.Find(&bookLogs)

	return bookLogs, result.Error
}
