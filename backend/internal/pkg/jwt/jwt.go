package jwt

import (
	"myapp/internal/exception"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/xerrors"
)

func GenerateToken(userId int) (string, error) {
	secretKey := os.Getenv("JWT_KEY")
	if secretKey == "" {
		return "", exception.ServerError
	}

	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", xerrors.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

func GetUserIDFromToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, xerrors.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return 0, xerrors.Errorf("failed to parse token: %w", err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["user_id"].(float64)

		return int(userID), nil
	} else {
		return 0, err
	}
}
