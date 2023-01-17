package models

import "time"

type BookLog struct {
	BookId    uint      `json:"book_id"`
	UserId    uint      `json:"user_id" gorm:"index"`
	LogType   string    `json:"log_type"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	User      User      `gorm:"foreignKey:UserId"`
	Book      Book      `gorm:"foreignKey:BookId"`
}
