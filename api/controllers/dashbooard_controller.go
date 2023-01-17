package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DashboardController struct{}

func (dc *DashboardController) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}
