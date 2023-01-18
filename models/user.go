package models

import "time"

const ADMIN_ROLE = "Admin"
const USER_ROLE = "User"

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"uniqueIndex;not null"`
	Role      string    `gorm:"type:varchar(255);not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
}

type UserRepository interface {
	Create(user *User) error
	GetByEmail(email string) (User, error)
	GetByID(id uint) (User, error)
}

func (u *User) isAdmin() bool {
	return u.Role == ADMIN_ROLE
}

func (u *User) isUser() bool {
	return u.Role == USER_ROLE
}
