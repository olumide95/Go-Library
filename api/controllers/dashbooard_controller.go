package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olumide95/go-library/domain"
)

type DashboardController struct {
	AuthUsecase domain.AuthUsecase
}

func (dc *DashboardController) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

func (dc *DashboardController) AdminHome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_dashboard.tmpl", gin.H{})
}

func (dc *DashboardController) UserHome(c *gin.Context) {
	c.HTML(http.StatusOK, "user_dashboard.tmpl", gin.H{})
}
