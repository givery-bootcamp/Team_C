package middleware

import (
	"errors"
	"net/http"
	"strings"

	"myapp/internal/config"
	"myapp/internal/exception"
	"myapp/internal/pkg/jwt"

	"github.com/gin-gonic/gin"
)

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

func GetUserIDFromContext(ctx *gin.Context) (int, error) {
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

	// Hostからポートを削除
	host := strings.Split(ctx.Request.Host, ":")[0]

	ctx.SetCookie(config.JWTCookieKeyName, token, 0, "/", host, false, false)
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

	// Hostからポートを削除
	host := strings.Split(ctx.Request.Host, ":")[0]

	ctx.SetCookie(config.JWTCookieKeyName, c, -1, "/", host, false, false)
	return nil
}
