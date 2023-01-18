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

	//Routes With DB Transaction
	txHandle := DB.Begin()
	brWithTx := repository.NewBookRepository(txHandle)
	blrWithTx := repository.NewBookLogRepository(txHandle)
	dcWithTx := controller.BookController{
		BookUsecase: usecase.NewbookUsecase(brWithTx, blrWithTx),
	}
	r.PUT("/books/borrow", middleware.DBTransactionMiddleware(txHandle), dcWithTx.BorrowBook)
	r.DELETE("/books/delete", middleware.DBTransactionMiddleware(txHandle), dc.DeleteBooks)
	r.PUT("/books/return", middleware.DBTransactionMiddleware(txHandle), dcWithTx.ReturnBook)
}
