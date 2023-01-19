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

func (ac *AuthController) Signup(c *gin.Context) {
	var request *domain.SignupRequest

	err := c.ShouldBind(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = ac.AuthUsecase.GetUserByEmail(request.Email)
	if err == nil {
		c.JSON(http.StatusNotFound, util.ErrorResponse{Message: "User already exists with the given email"})
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
		Role:     models.USER_ROLE,
	}

	err = ac.AuthUsecase.CreateUser(&user)
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

	access_token, _ := util.CreateToken(user.Email)

	c.SetCookie("access_token", access_token, 86400, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", 86400, "/", "localhost", false, false)

	c.JSON(http.StatusOK, gin.H{"user": userResponse, "access_token": access_token})
}

func (ac *AuthController) Login(c *gin.Context) {
	var request *domain.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := ac.AuthUsecase.GetUserByEmail(request.Email)

	if err != nil {
		c.JSON(http.StatusNotFound, util.ErrorResponse{Message: "User not found with the given email"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, util.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	loginResponse := &domain.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}

	access_token, _ := util.CreateToken(user.Email)

	c.SetCookie("access_token", access_token, 86400, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", 86400, "/", "localhost", false, false)

	c.JSON(http.StatusOK, gin.H{"user": loginResponse, "access_token": access_token})
}

func (ac *AuthController) LoginView(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.tmpl", gin.H{})
}

func (ac *AuthController) SignupView(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.tmpl", gin.H{})
}

func (ac *AuthController) LogoutUser(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("logged_in", "", -1, "/", "localhost", false, false)

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
