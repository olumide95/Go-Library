package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/olumide95/go-library/api/controllers"
	"github.com/olumide95/go-library/repository"
	"github.com/olumide95/go-library/usecase"
	"gorm.io/gorm"
)

func BookRouter(r *gin.RouterGroup, DB *gorm.DB) {
	ur := repository.NewBookRepository(DB)
	dc := controller.BookController{
		BookUsecase: usecase.NewbookUsecase(ur),
	}
	r.GET("/books/all", dc.AllBooks)
	r.POST("/books/store", dc.StoreBooks)
	r.POST("/books/update", dc.UpdateBook)
}
