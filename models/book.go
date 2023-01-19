package models

import "time"

type Book struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Quantity  uint16    `json:"quantity"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
}

type BookRepository interface {
	Create(book *Book) error
	Update(id uint, book *Book) (int64, error)
	CreateBulk(books *[]Book) error
	GetByID(id uint) (Book, error)
	GetByIds(ids []uint) ([]Book, error)
	GetByIDForUpdate(id uint) (Book, error)
	Delete(id []uint) (int64, error)
	All() ([]Book, error)
}
