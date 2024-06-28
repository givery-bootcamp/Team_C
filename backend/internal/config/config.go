package config

import (
	"os"
	"strconv"
)

var (
	HostName         = "127.0.0.1"
	DomainURL        = "localhost"
	Port             = 9000
	CorsAllowOrigin  = []string{"http://localhost:3000", "http://localhost:8001"}
	DBHostName       = "db"
	DBUser           = "root"
	DBPassword       = "password"
	DBPort           = 3306
	DBName           = "training"
	JWTCookieKeyName = "Authorize"
	GinSigninUserKey = "userID"
)

func init() {
	LoadConfig()
}

func LoadConfig() {
	if v := os.Getenv("HOSTNAME"); v != "" {
		HostName = v
	}
	if v := os.Getenv("DOMAIN_URL"); v != "" {
		DomainURL = v
	}
	if v, err := strconv.ParseInt(os.Getenv("PORT"), 10, 64); err == nil {
		Port = int(v)
	}
	if v := os.Getenv("CORS_ALLOW_ORIGIN"); v != "" {
		CorsAllowOrigin = append(CorsAllowOrigin, v)
	}
	if v := os.Getenv("DB_HOSTNAME"); v != "" {
		DBHostName = v
	}
	if v := os.Getenv("DB_USERNAME"); v != "" {
		DBUser = v
	}
	if v := os.Getenv("DB_PASSWORD"); v != "" {
		DBPassword = v
	}
	if v, err := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64); err == nil {
		DBPort = int(v)
	}
	if v := os.Getenv("DB_NAME"); v != "" {
		DBName = v
	}
}
