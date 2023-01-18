package models

import "time"

type BookLog struct {
	ID         uint       `json:"id" gorm:"primary_key"`
	BookId     uint       `json:"book_id"`
	UserId     uint       `json:"user_id" gorm:"index"`
	BorrowedAt *time.Time `gorm:"default:current_timestamp"`
	ReturnedAt *time.Time `gorm:"default:null"`
	User       User       `gorm:"foreignKey:UserId"`
	Book       Book       `gorm:"foreignKey:BookId"`
}

type BookLogRepository interface {
	Create(bookLog *BookLog) error
	Update(id uint, bookLog *BookLog) (int64, error)
	GetForUpdate(id uint, userId uint) (BookLog, error)
	Delete(id []uint) error
	All() ([]BookLog, error)
}
