package controller

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/olumide95/go-library/api/util"
	"github.com/olumide95/go-library/domain"
)

type DashboardController struct {
	AuthUsecase domain.AuthUsecase
}

func (dc *DashboardController) Home(c *gin.Context) {
	session := util.NewSession(c)
	userEmail := session.GetSessionData(os.Getenv("USER_SESSION_KEY"))

	c.HTML(http.StatusOK, "dashboard.tmpl", gin.H{"email": userEmail})
}
