package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/olumide95/go-library/api/controllers"
	"github.com/olumide95/go-library/repository"
	"github.com/olumide95/go-library/usecase"
	"gorm.io/gorm"
)

func DashboardRouter(r *gin.RouterGroup, DB *gorm.DB) {
	ur := repository.NewUserRepository(DB)
	dc := controller.DashboardController{
		AuthUsecase: usecase.NewauthUsecase(ur),
	}
	r.GET("/admin", dc.AdminHome)
	r.GET("/user", dc.UserHome)
}
