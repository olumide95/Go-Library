package repository

import (
	"strings"

	"github.com/olumide95/go-library/models"
	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(DB *gorm.DB) models.UserRepository {
	return &userRepository{DB}
}

func (ur *userRepository) Create(user *models.User) error {

	result := ur.database.Create(user)

	return result.Error
}

func (ur *userRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	result := ur.database.First(&user, "email = ?", strings.ToLower(email))

	return user, result.Error
}

func (ur *userRepository) GetByID(id uint) (models.User, error) {
	var user models.User
	result := ur.database.First(&user, "id = ?", id)

	return user, result.Error
}
