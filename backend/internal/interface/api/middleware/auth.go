//go:generate mockgen -source=auth.go -destination=middleware_mock/auth_mock.go -package middleware_mock
package middleware

import (
	"errors"
	"net/http"

	"myapp/internal/config"
	"myapp/internal/exception"
	"myapp/internal/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	CheckToken() gin.HandlerFunc
	GetUserIDFromContext(ctx *gin.Context) (int, error)
	SetJWTCookie(ctx *gin.Context, userID int) error
	DeleteCookie(ctx *gin.Context) error
}

func CheckToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie(config.JWTCookieKeyName)
		if err != nil {
			ctx.Error(exception.AuthError)
			ctx.Abort()
			return
		}

		userID, err := jwt.GetUserIDFromToken(token)
		if err != nil {
			err, ok := err.(*exception.Exception)
			if ok {
				ctx.Error(err)
			} else {
				ctx.Error(exception.AuthError)
			}
			ctx.Abort()
			return
		}

		ctx.Set(config.GinSigninUserKey, userID)
		ctx.Next()
	}
}

var GetUserIDFromContext = func(ctx *gin.Context) (int, error) {
	u, isExists := ctx.Get(config.GinSigninUserKey)
	if !isExists {
		return 0, exception.ServerError
	}

	return u.(int), nil
}

func SetJWTCookie(ctx *gin.Context, userID int) error {
	token, err := jwt.GenerateToken(userID)
	if err != nil {
		return exception.ServerError
	}

	ctx.SetCookie(config.JWTCookieKeyName, token, 0, "/", "localhost", false, true)

	return nil
}

func DeleteCookie(ctx *gin.Context) error {
	c, err := ctx.Cookie(config.JWTCookieKeyName)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return nil
		}
		return err
	}

	ctx.SetCookie(config.JWTCookieKeyName, c, -1, "/", ctx.Request.Host, false, true)
	return nil
}
