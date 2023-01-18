package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olumide95/go-library/api/util"
	"github.com/olumide95/go-library/domain"
	"github.com/olumide95/go-library/models"
	csrf "github.com/utrack/gin-csrf"
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

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}

func (ac *AuthController) Login(c *gin.Context) {
	var request *domain.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := ac.AuthUsecase.GetUserByEmail(request.Email)

	session := util.NewSession(c)

	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		session.SetFlashMessage("User not found with the given email.")
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		session.SetFlashMessage("Invalid credentials.")
		return
	}

	loginResponse := &domain.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}

	c.JSON(http.StatusOK, gin.H{"user": loginResponse})
}

func (ac *AuthController) LoginView(c *gin.Context) {
	session := util.NewSession(c)
	c.HTML(http.StatusOK, "signin.tmpl", gin.H{"messages": session.GetFlashMessage(), "csrf": csrf.GetToken(c)})
}
