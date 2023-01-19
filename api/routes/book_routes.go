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
	r.GET("/books/borrowed-books", dc.AllBorrowedBooks)
}

func BookRouterWithTx(r *gin.RouterGroup, DB *gorm.DB) {
	br := repository.NewBookRepository(DB)
	blr := repository.NewBookLogRepository(DB)
	dc := controller.BookController{
		BookUsecase: usecase.NewbookUsecase(br, blr),
	}
	//Routes With DB Transaction
	r.PATCH("/books/update", middleware.DBTransactionMiddleware(DB), dc.UpdateBook)
	r.PATCH("/books/borrow", middleware.DBTransactionMiddleware(DB), dc.BorrowBook)
	r.DELETE("/books/delete", middleware.DBTransactionMiddleware(DB), dc.DeleteBooks)
	r.PATCH("/books/return", middleware.DBTransactionMiddleware(DB), dc.ReturnBook)
}
