package routes

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/olumide95/go-library/api/middleware"
	"gorm.io/gorm"
)

func Setup(DB *gorm.DB) {
	router := gin.Default()
	secret := os.Getenv("SESSION_SECRET")

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/**/*.tmpl")

	router.Use(sessions.Sessions("session", cookie.NewStore([]byte(secret))))

	PublicRoutes(router.Group("/"), DB)

	private := router.Group("/")
	private.Use(middleware.NewAuthenticatedMiddlware(DB).Check)
	PrivateRoutes(private, DB)

	router.Run()
}

func PublicRoutes(router *gin.RouterGroup, DB *gorm.DB) {
	AuthRouter(router, DB)
	DashboardRouter(router, DB)
	BookRouter(router, DB)
}

func PrivateRoutes(router *gin.RouterGroup, DB *gorm.DB) {

}
