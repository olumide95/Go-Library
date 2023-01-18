package controller

import (
	"net/http"
	"os"

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

	session := util.NewSession(c)

	if err != nil {
		c.Redirect(http.StatusSeeOther, "/signup")
		session.SetFlashMessage(err.Error())
		return
	}

	_, err = ac.AuthUsecase.GetUserByEmail(request.Email)
	if err == nil {
		c.Redirect(http.StatusSeeOther, "/signup")
		session.SetFlashMessage("User already exists with the given email")
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		c.Redirect(http.StatusSeeOther, "/signup")
		session.SetFlashMessage(err.Error())
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
		c.Redirect(http.StatusSeeOther, "/signup")
		session.SetFlashMessage(err.Error())
		return
	}

	session.SetSessionData(os.Getenv("USER_SESSION_KEY"), user.Email)
	c.Redirect(http.StatusSeeOther, "/")
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

	session.SetSessionData(os.Getenv("USER_SESSION_KEY"), user.Email)
	c.Redirect(http.StatusSeeOther, "/")
}

func (ac *AuthController) LoginView(c *gin.Context) {
	session := util.NewSession(c)
	c.HTML(http.StatusOK, "signin.tmpl", gin.H{"messages": session.GetFlashMessage(), "csrf": csrf.GetToken(c)})
}

func (ac *AuthController) SignupView(c *gin.Context) {
	session := util.NewSession(c)
	c.HTML(http.StatusOK, "signup.tmpl", gin.H{"messages": session.GetFlashMessage(), "csrf": csrf.GetToken(c)})
}
