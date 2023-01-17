package routes

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/olumide95/go-library/api/middleware"
	csrf "github.com/utrack/gin-csrf"
	"gorm.io/gorm"
)

func Setup(DB *gorm.DB) {
	router := gin.Default()
	secret := os.Getenv("SESSION_SECRET")

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/**/*.tmpl")

	router.Use(sessions.Sessions("session", cookie.NewStore([]byte(secret))))
	router.Use(csrf.Middleware(csrf.Options{
		Secret: os.Getenv("CSRF_SECRET"),
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))

	PublicRoutes(router.Group("/"), DB)

	private := router.Group("/")
	private.Use(middleware.AuthRequired)
	PrivateRoutes(private, DB)

	router.Run()
}

func PublicRoutes(router *gin.RouterGroup, DB *gorm.DB) {
	AuthRouter(router, DB)
}

func PrivateRoutes(router *gin.RouterGroup, DB *gorm.DB) {
	DashboardRouter(router, DB)
}
