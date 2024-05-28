package driver

import (
	"fmt"
	"myapp/internal/config"
)

var dsn = fmt.Sprintf("%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBHostName, config.DBPort, config.DBName)

const dsnPrefixForGomigrate = "mysql://"

func createDSNForGoMigrate() string {
	return fmt.Sprintf("%s%s", dsnPrefixForGomigrate, dsn)
}

func createDSNForGorm() string {
	return dsn
}
