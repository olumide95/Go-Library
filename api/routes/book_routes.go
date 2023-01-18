package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/olumide95/go-library/api/controllers"
	"github.com/olumide95/go-library/api/middleware"
	"github.com/olumide95/go-library/repository"
	"github.com/olumide95/go-library/usecase"
	"gorm.io/gorm"
)

func BookRouter(r *gin.RouterGroup, DB *gorm.DB) {
	br := repository.NewBookRepository(DB)
	blr := repository.NewBookLogRepository(DB)
	dc := controller.BookController{
		BookUsecase: usecase.NewbookUsecase(br, blr),
	}

	r.GET("/books/all", dc.AllBooks)
	r.POST("/books/store", dc.StoreBooks)
	r.PUT("/books/update", dc.UpdateBook)
	r.DELETE("/books/delete", dc.DeleteBooks)
	r.PUT("/books/borrow", middleware.DBTransactionMiddleware(DB), dc.BorrowBook)
	r.PUT("/books/return", middleware.DBTransactionMiddleware(DB), dc.ReturnBook)
}
