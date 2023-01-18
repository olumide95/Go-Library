package models

import "time"

type BookLog struct {
	BookId     uint      `json:"book_id"`
	UserId     uint      `json:"user_id" gorm:"index"`
	BorrowedAt time.Time `gorm:"default:current_timestamp"`
	ReturnedAt time.Time `json:"returned_at"`
	User       User      `gorm:"foreignKey:UserId"`
	Book       Book      `gorm:"foreignKey:BookId"`
}
