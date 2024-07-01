package jwt_test

import (
	"myapp/internal/exception"
	"myapp/internal/pkg/jwt"
	"testing"
	"time"

	jwt_lib "github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func setupEnv(t *testing.T, envs map[string]string) {
	for k, v := range envs {
		t.Setenv(k, v)
	}
}

func TestGenerateToken(t *testing.T) {
	t.Run("JWT_KEYが空の場合エラーが返る", func(t *testing.T) {
		setupEnv(t, map[string]string{"JWT_KEY": ""})
		token, err := jwt.GenerateToken(123)
		assert.Error(t, err)
		assert.Equal(t, exception.ServerError, err)
		assert.Empty(t, token)
	})

	t.Run("Success", func(t *testing.T) {
		setupEnv(t, map[string]string{"JWT_KEY": "test_secret_key"})
		token, err := jwt.GenerateToken(123)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})
}

func TestGetUserIDFromToken(t *testing.T) {
	setupEnv(t, map[string]string{"JWT_KEY": "test_secret_key"})

	t.Run("Tokenが適切な形式でない場合エラーを返す", func(t *testing.T) {
		userID, err := jwt.GetUserIDFromToken("invalid_token")
		assert.Equal(t, "failed to parse token: token is malformed: token contains an invalid number of segments", err.Error())
		assert.Equal(t, 0, userID)
	})

	t.Run("Tokenが期限切れの場合エラーを返す", func(t *testing.T) {
		token := jwt_lib.NewWithClaims(jwt_lib.SigningMethodHS256, jwt_lib.MapClaims{
			"user_id": float64(123),
			"exp":     time.Now().Add(-time.Hour).Unix(), // 過去の時間
		})
		tokenString, _ := token.SignedString([]byte("test_secret_key"))

		userID, err := jwt.GetUserIDFromToken(tokenString)
		assert.Equal(t, "failed to parse token: token has invalid claims: token is expired", err.Error())
		assert.Equal(t, 0, userID)
	})

	t.Run("署名方法が異なる場合エラーを返す", func(t *testing.T) {
		token := jwt_lib.NewWithClaims(jwt_lib.SigningMethodNone, jwt_lib.MapClaims{
			"user_id": float64(123),
			"exp":     time.Now().Add(time.Hour).Unix(),
		})
		tokenString, _ := token.SignedString(jwt_lib.UnsafeAllowNoneSignatureType)

		userID, err := jwt.GetUserIDFromToken(tokenString)
		assert.Equal(t, "failed to parse token: token is unverifiable: error while executing keyfunc: Unexpected signing method: none", err.Error())
		assert.Equal(t, 0, userID)
	})

	t.Run("Claimがない場合エラーを返す", func(t *testing.T) {
		token := jwt_lib.NewWithClaims(jwt_lib.SigningMethodHS256, jwt_lib.MapClaims{})
		tokenString, _ := token.SignedString([]byte("test_secret_key"))

		userID, err := jwt.GetUserIDFromToken(tokenString)
		assert.Equal(t, "token invalid", err.Error())
		assert.Equal(t, 0, userID)
	})

	t.Run("success", func(t *testing.T) {
		token := jwt_lib.NewWithClaims(jwt_lib.SigningMethodHS256, jwt_lib.MapClaims{
			"user_id": float64(123),
			"exp":     time.Now().Add(time.Hour).Unix(),
		})
		tokenString, _ := token.SignedString([]byte("test_secret_key"))

		userID, err := jwt.GetUserIDFromToken(tokenString)
		assert.NoError(t, err)
		assert.Equal(t, 123, userID)
	})
}
