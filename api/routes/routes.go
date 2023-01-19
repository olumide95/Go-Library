package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/olumide95/go-library/api/middleware"
	"gorm.io/gorm"
)

func Setup(DB *gorm.DB) {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/**/*.tmpl")

	PublicRoutes(router.Group("/"), DB)

	private := router.Group("/")
	private.Use(middleware.NewAuthenticatedMiddlware(DB).Check)
	PrivateRoutes(private, DB)

	router.Run()
}

func PublicRoutes(router *gin.RouterGroup, DB *gorm.DB) {
	AuthRouter(router, DB)
	DashboardRouter(router, DB)
}

func PrivateRoutes(router *gin.RouterGroup, DB *gorm.DB) {
	BookRouter(router, DB)
	BookRouterWithTx(router, DB)
}
