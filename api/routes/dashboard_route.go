package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/olumide95/go-library/api/controllers"
	"gorm.io/gorm"
)

func DashboardRouter(r *gin.RouterGroup, DB *gorm.DB) {
	dc := controller.DashboardController{}
	r.GET("/", dc.Home)
}
