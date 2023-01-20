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

func (dc *DashboardController) BorrowBook(c *gin.Context) {
	c.HTML(http.StatusOK, "user_borrow_book.tmpl", gin.H{})
}

func (dc *DashboardController) ReturnBook(c *gin.Context) {
	c.HTML(http.StatusOK, "user_return_book.tmpl", gin.H{})
}

func (dc *DashboardController) AddBooks(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_add_books.tmpl", gin.H{})
}

func (dc *DashboardController) ViewBooks(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_view_books.tmpl", gin.H{})
}
