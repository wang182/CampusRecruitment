package middleware

import (
	"CampusRecruitment/pkg/apps/ctx"
	"CampusRecruitment/pkg/config"
	"CampusRecruitment/pkg/types"
	"CampusRecruitment/pkg/types/errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ac := ctx.New(c)
		tokenStr := ac.Token()
		if tokenStr == "" {
			ac.Response(nil, errors.ErrNotAuth.WithStatus(http.StatusUnauthorized))
			c.Abort()
			return
		}

		claims := types.UserTokenClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Get().SecretKey), nil
		})

		if err != nil || token == nil || !token.Valid {
			ac.Response(nil, errors.ErrNotAuth.WithCause(err).WithStatus(http.StatusUnauthorized))
			c.Abort()
			return
		}
		c.Next()
	}
}
