package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/olumide95/go-library/api/util"
	"github.com/olumide95/go-library/domain"
	"github.com/olumide95/go-library/repository"
	"github.com/olumide95/go-library/usecase"
	"gorm.io/gorm"

	"net/http"
)

type AuthenticatedMiddlware struct {
	AuthUsecase domain.AuthUsecase
}

func NewAuthenticatedMiddlware(DB *gorm.DB) *AuthenticatedMiddlware {
	ur := repository.NewUserRepository(DB)
	return &AuthenticatedMiddlware{
		AuthUsecase: usecase.NewauthUsecase(ur),
	}
}

func (ar *AuthenticatedMiddlware) Check(c *gin.Context) {
	var accessToken string
	cookie, err := c.Cookie("accessToken")

	authorizationHeader := c.Request.Header.Get("Authorization")
	fields := strings.Fields(authorizationHeader)

	if len(fields) != 0 && fields[0] == "Bearer" {
		accessToken = fields[1]
	} else if err == nil {
		accessToken = cookie
	}

	if accessToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, util.ErrorResponse{Message: "You are not logged in"})
		return
	}

	sub, err := util.ValidateToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, util.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := ar.AuthUsecase.GetUserByEmail(fmt.Sprint(sub))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, util.ErrorResponse{Message: "User does not exist"})
		return
	}

	c.Set("currentUserId", user.ID)
	c.Next()
}
