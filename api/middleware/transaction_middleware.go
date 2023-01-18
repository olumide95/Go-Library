package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StatusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		if i == status {
			return true
		}
	}
	return false
}

func DBTransactionMiddleware(txHandle *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Print("beginning database transaction")

		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
			}
		}()

		c.Next()

		if StatusInList(c.Writer.Status(), []int{http.StatusOK, http.StatusCreated}) {
			log.Print("committing transactions")
			if err := txHandle.Commit().Error; err != nil {
				log.Print("trx commit error: ", err)
			}
		} else {
			log.Print("rolling back transaction due to status code: ", c.Writer.Status())
			txHandle.Rollback()
		}
	}
}
