package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"myapp/internal/config"
	"myapp/internal/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupEnv(t *testing.T, envs map[string]string) {
	for k, v := range envs {
		t.Setenv(k, v)
	}
}

func TestCheckToken(t *testing.T) {
	t.Run("failed/Cookieセットされていない場合エラーを返す", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/", nil)

		middleware := CheckToken()
		middleware(c)

		assert.NotNil(t, c.Errors.Last())
	})

	t.Run("failed/Tokenが不正な場合エラーを返す", func(t *testing.T) {
		setupEnv(t, map[string]string{"JWT_KEY": "test_secret_key"})

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/", nil)
		c.Request.AddCookie(&http.Cookie{Name: config.JWTCookieKeyName, Value: "invalid_token"})

		middleware := CheckToken()
		middleware(c)

		assert.NotNil(t, c.Errors.Last())
	})

	t.Run("success", func(t *testing.T) {
		setupEnv(t, map[string]string{"JWT_KEY": "test_secret_key"})

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/", nil)

		// JWT生成
		token, err := jwt.GenerateToken(1)
		require.NoError(t, err)
		c.Request.AddCookie(&http.Cookie{Name: config.JWTCookieKeyName, Value: token})

		middleware := CheckToken()
		middleware(c)

		assert.Nil(t, c.Errors.Last())

		userID, exists := c.Get(config.GinSigninUserKey)
		assert.True(t, exists)
		assert.Equal(t, 1, userID)
	})
}

func TestGetUserIDFromContext(t *testing.T) {
	t.Run("failed/ContextにUserIDがセットされていない場合エラーを返す", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		_, err := GetUserIDFromContext(c)

		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set(config.GinSigninUserKey, 1)

		userID, err := GetUserIDFromContext(c)
		assert.NoError(t, err)
		assert.Equal(t, 1, userID)
	})
}

func TestSetJWTCookie(t *testing.T) {
	t.Run("failed/Token生成に失敗した場合エラーを返す", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		// JWT_KEYが未設定なのでエラーが返る
		err := SetJWTCookie(c, 1)

		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		setupEnv(t, map[string]string{"JWT_KEY": "test_secret_key"})

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		err := SetJWTCookie(c, 1)

		assert.NoError(t, err)
		assert.NotEmpty(t, w.Result().Cookies())
	})
}

func TestDeleteCookie(t *testing.T) {
	t.Run("success/Cookieが存在しない場合何もしない", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/", nil)
		c.Request.AddCookie(&http.Cookie{Name: "dummy", Value: "dummy"})

		DeleteCookie(c)
	})

	t.Run("success", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/", nil)
		c.Request.AddCookie(&http.Cookie{Name: config.JWTCookieKeyName, Value: "test_token"})

		DeleteCookie(c)
		assert.NotEmpty(t, w.Result().Cookies())
	})
}
