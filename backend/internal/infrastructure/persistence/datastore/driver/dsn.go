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
	dbpass := config.DBPassword
	return fmt.Sprintf("mysql://%s:%s@tcp(%s:%d)/%s%s", dbuser, dbpass, host, port, dbname, DSNSuffix)
}

func createDSNForGorm() string {
	host := config.DBHostName
	port := config.DBPort
	dbname := config.DBName
	dbuser := config.DBUser
	dbpass := config.DBPassword
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s", dbuser, dbpass, host, port, dbname, DSNSuffix)
}
