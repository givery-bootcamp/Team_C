package middleware

import (
	"github.com/gin-gonic/gin"
)

func SetCookie(ctx *gin.Context, token string) error {
	ctx.SetCookie("Authorise", token, 0, "/", ctx.Request.Host, false, true)

	return nil
}
