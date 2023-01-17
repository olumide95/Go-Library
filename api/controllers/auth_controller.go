package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olumide95/go-library/api/util"
	"github.com/olumide95/go-library/domain"
	"github.com/olumide95/go-library/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	AuthUsecase domain.AuthUsecase
}

func (sc *AuthController) Signup(c *gin.Context) {
	var request *domain.SignupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = sc.AuthUsecase.GetUserByEmail(request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, util.ErrorResponse{Message: "User already exists with the given email"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ErrorResponse{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = sc.AuthUsecase.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ErrorResponse{Message: err.Error()})
		return
	}

	userResponse := &domain.SignupResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}
