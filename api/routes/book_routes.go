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
	bc := controller.BookController{
		BookUsecase: usecase.NewbookUsecase(br, blr),
	}

	r.GET("/books/all", bc.AllBooks)
	r.POST("/books/store", bc.StoreBooks)
	r.GET("/books/borrowed-books", bc.AllBorrowedBooks)
	r.GET("/books/get-book", bc.GetBook)
}

func BookRouterWithTx(r *gin.RouterGroup, DB *gorm.DB) {
	br := repository.NewBookRepository(DB)
	blr := repository.NewBookLogRepository(DB)
	bc := controller.BookController{
		BookUsecase: usecase.NewbookUsecase(br, blr),
	}
	//Routes With DB Transaction
	r.PATCH("/books/update", middleware.DBTransactionMiddleware(DB), bc.UpdateBook)
	r.PATCH("/books/borrow", middleware.DBTransactionMiddleware(DB), bc.BorrowBook)
	r.DELETE("/books/delete", middleware.DBTransactionMiddleware(DB), bc.DeleteBooks)
	r.PATCH("/books/return", middleware.DBTransactionMiddleware(DB), bc.ReturnBook)
}
