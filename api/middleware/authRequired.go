package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/olumide95/go-library/api/util"

	"log"
	"net/http"
)

func AuthRequired(c *gin.Context) {
	session := util.NewSession(c)
	user := session.GetSessionData(os.Getenv("USER_SESSION_KEY"))
	if user == nil {
		log.Println("User not logged in")
		c.Redirect(http.StatusMovedPermanently, "/login")
		c.Abort()
		return
	}
	c.Next()
}
