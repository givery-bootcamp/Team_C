package middleware

import (
	myConfig "myapp/internal/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = myConfig.CorsAllowOrigin
	config.AllowCredentials = true
	config.AllowHeaders = []string{
		"Access-Control-Allow-Credentials",
		"Access-Control-Allow-Headers",
		"Content-Type",
		"Content-Length",
		"Authorization",
	}
	return cors.New(config)
}
