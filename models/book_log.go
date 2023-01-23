package models

import (
	"time"

	"gorm.io/gorm"
)

type BookLog struct {
	ID         uint       `json:"id" gorm:"primary_key"`
	BookId     uint       `json:"book_id"`
	UserId     uint       `json:"user_id" gorm:"index"`
	BorrowedAt *time.Time `gorm:"default:current_timestamp"`
	ReturnedAt *time.Time `gorm:"default:null"`
	User       User       `json:"-" gorm:"foreignKey:UserId"`
	Book       Book       `gorm:"foreignKey:BookId"`
}

type BookLogRepository interface {
	Create(bookLog *BookLog) error
	GetForUpdate(id uint, userId uint) (BookLog, error)
	GetWithBooks(userId uint) ([]BookLog, error)
	Update(id uint, bookLog *BookLog) (int64, error)
	Delete(id []uint) error
	DeleteByBookIds(bookIds []uint) error
	All() ([]BookLog, error)
	WithTrx(txHandle *gorm.DB) BookLogRepository
}
