package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/olumide95/go-library/api/controllers"
	"github.com/olumide95/go-library/repository"
	"github.com/olumide95/go-library/usecase"
	"gorm.io/gorm"
)

func AuthRouter(r *gin.Engine, DB *gorm.DB) {
	ur := repository.NewUserRepository(DB)
	sc := controller.AuthController{
		AuthUsecase: usecase.NewauthUsecase(ur),
	}
	r.POST("/signup", sc.Signup)
}
