package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	myConfig "myapp/internal/config"
	"myapp/internal/interface/api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCors(t *testing.T) {
	gin.SetMode(gin.TestMode)
	myConfig.CorsAllowOrigin = []string{"http://example.com"}

	// Ginルーターの設定
	router := gin.New()
	router.Use(middleware.Cors())
	router.GET("/test", func(c *gin.Context) {
		c.String(200, "OK")
	})

	req, _ := http.NewRequest("OPTIONS", "/test", nil)
	req.Header.Set("Origin", "http://example.com")
	req.Header.Set("Access-Control-Request-Method", "GET")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, "http://example.com", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "true", w.Header().Get("Access-Control-Allow-Credentials"))

	assert.Contains(t, w.Header().Get("Access-Control-Allow-Headers"), "Content-Length")
	assert.Contains(t, w.Header().Get("Access-Control-Allow-Headers"), "Content-Type")
	assert.Contains(t, w.Header().Get("Access-Control-Allow-Headers"), "Access-Control-Allow-Headers")
	assert.Contains(t, w.Header().Get("Access-Control-Allow-Headers"), "Access-Control-Allow-Credentials")
	assert.Contains(t, w.Header().Get("Access-Control-Allow-Headers"), "Authorization")
}
