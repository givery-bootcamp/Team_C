package driver

import (
	"fmt"
	"myapp/internal/config"
)

const DSNSuffix = "?charset=utf8mb4&parseTime=True&loc=Local"

func createDSNForGoMigrate() string {
	host := config.DBHostName
	port := config.DBPort
	dbname := config.DBName
	dbuser := config.DBUser
	return fmt.Sprintf("mysql://%s@tcp(%s:%d)/%s%s", dbuser, host, port, dbname, DSNSuffix)
}

func createDSNForGorm() string {
	host := config.DBHostName
	port := config.DBPort
	dbname := config.DBName
	dbuser := config.DBUser

	return fmt.Sprintf("%s@tcp(%s:%d)/%s%s", dbuser, host, port, dbname, DSNSuffix)
}
