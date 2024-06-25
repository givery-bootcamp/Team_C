package jwt_test

import (
	"myapp/internal/exception"
	"myapp/internal/pkg/jwt"
	"os"
	"testing"
	"time"

	jwt_lib "github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	// 元の環境変数を保存
	originalJWTKey := os.Getenv("JWT_KEY")
	defer os.Setenv("JWT_KEY", originalJWTKey)

	t.Run("Success", func(t *testing.T) {
		os.Setenv("JWT_KEY", "test_secret_key")
		token, err := jwt.GenerateToken(123)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})

	t.Run("EmptySecretKey", func(t *testing.T) {
		os.Setenv("JWT_KEY", "")
		token, err := jwt.GenerateToken(123)
		assert.Error(t, err)
		assert.Equal(t, exception.ServerError, err)
		assert.Empty(t, token)
	})
}

func TestGetUserIDFromToken(t *testing.T) {
	// 元の環境変数を保存
	originalJWTKey := os.Getenv("JWT_KEY")
	defer os.Setenv("JWT_KEY", originalJWTKey)

	os.Setenv("JWT_KEY", "test_secret_key")

	t.Run("ValidToken", func(t *testing.T) {
		token := jwt_lib.NewWithClaims(jwt_lib.SigningMethodHS256, jwt_lib.MapClaims{
			"user_id": float64(123),
			"exp":     time.Now().Add(time.Hour).Unix(),
		})
		tokenString, _ := token.SignedString([]byte("test_secret_key"))

		userID, err := jwt.GetUserIDFromToken(tokenString)
		assert.NoError(t, err)
		assert.Equal(t, 123, userID)
	})

	t.Run("InvalidToken", func(t *testing.T) {
		userID, err := jwt.GetUserIDFromToken("invalid_token")
		assert.Error(t, err)
		assert.Equal(t, 0, userID)
	})

	t.Run("ExpiredToken", func(t *testing.T) {
		token := jwt_lib.NewWithClaims(jwt_lib.SigningMethodHS256, jwt_lib.MapClaims{
			"user_id": float64(123),
			"exp":     time.Now().Add(-time.Hour).Unix(), // 過去の時間
		})
		tokenString, _ := token.SignedString([]byte("test_secret_key"))

		userID, err := jwt.GetUserIDFromToken(tokenString)
		assert.Error(t, err)
		assert.Equal(t, 0, userID)
	})

	t.Run("WrongSigningMethod", func(t *testing.T) {
		token := jwt_lib.NewWithClaims(jwt_lib.SigningMethodNone, jwt_lib.MapClaims{
			"user_id": float64(123),
			"exp":     time.Now().Add(time.Hour).Unix(),
		})
		tokenString, _ := token.SignedString(jwt_lib.UnsafeAllowNoneSignatureType)

		userID, err := jwt.GetUserIDFromToken(tokenString)
		assert.Error(t, err)
		assert.Equal(t, 0, userID)
	})
}
