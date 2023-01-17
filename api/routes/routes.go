package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(DB *gorm.DB) {
	r := gin.Default()

	AuthRouter(r, DB)

	r.Run()
}
