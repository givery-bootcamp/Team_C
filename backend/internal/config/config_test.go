package config_test

import (
	"myapp/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupEnv(t *testing.T, envs map[string]string) {
	for k, v := range envs {
		t.Setenv(k, v)
	}
}

func TestConfigInitialization(t *testing.T) {
	testCases := []struct {
		name     string
		envVars  map[string]string
		expected map[string]interface{}
	}{
		{
			name:    "success/デフォルト値",
			envVars: map[string]string{},
			expected: map[string]interface{}{
				"HostName":   "127.0.0.1",
				"Port":       9000,
				"DBHostName": "db",
				"DBUser":     "root",
				"DBPassword": "password",
				"DBPort":     3306,
				"DBName":     "training",
				"DomainURL":  "localhost",
			},
		},
		{
			name: "success/環境変数で上書き",
			envVars: map[string]string{
				"HOSTNAME":          "0.0.0.0",
				"PORT":              "8080",
				"CORS_ALLOW_ORIGIN": "http://example.com",
				"DB_HOSTNAME":       "custom-db",
				"DB_USERNAME":       "admin",
				"DB_PASSWORD":       "secret",
				"DB_PORT":           "3307",
				"DB_NAME":           "test-db",
				"DomainURL":         "customdomain",
			},
			expected: map[string]interface{}{
				"HostName":   "0.0.0.0",
				"Port":       8080,
				"DBHostName": "custom-db",
				"DBUser":     "admin",
				"DBPassword": "secret",
				"DBPort":     3307,
				"DBName":     "test-db",
				"DomainURL":  "customdomain",
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			setupEnv(t, tt.envVars)

			config.LoadConfig()

			assert.Equal(t, tt.expected["HostName"], config.HostName)
			assert.Equal(t, tt.expected["Port"], config.Port)
			assert.Equal(t, tt.expected["DBHostName"], config.DBHostName)
			assert.Equal(t, tt.expected["DBUser"], config.DBUser)
			assert.Equal(t, tt.expected["DBPassword"], config.DBPassword)
			assert.Equal(t, tt.expected["DBPort"], config.DBPort)
			assert.Equal(t, tt.expected["DBName"], config.DBName)
			assert.Equal(t, tt.expected["DomainURL"], config.DomainURL)

			if v, ok := tt.envVars["CORS_ALLOW_ORIGIN"]; ok {
				assert.Contains(t, config.CorsAllowOrigin, v)
			}
		})
	}
}
