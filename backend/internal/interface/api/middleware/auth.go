package middleware

import (
	"myapp/internal/config"
	"myapp/internal/exception"
	"myapp/internal/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func SetJWTCookie(ctx *gin.Context, userID int) error {
	token, err := jwt.GenerateToken(userID)
	if err != nil {
		return exception.ServerError
	}
	ctx.SetCookie(config.JWTCookieKeyName, token, 0, "/", ctx.Request.Host, false, true)

	return nil
}
