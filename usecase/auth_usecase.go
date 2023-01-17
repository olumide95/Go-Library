package usecase

import (
	"github.com/olumide95/go-library/domain"
	"github.com/olumide95/go-library/models"
)

type authUsecase struct {
	userRepository models.UserRepository
}

func NewauthUsecase(userRepository models.UserRepository) domain.AuthUsecase {
	return &authUsecase{
		userRepository: userRepository,
	}
}

func (su *authUsecase) CreateUser(user *models.User) error {
	return su.userRepository.Create(user)
}

func (su *authUsecase) GetUserByEmail(email string) (models.User, error) {
	return su.userRepository.GetByEmail(email)
}
