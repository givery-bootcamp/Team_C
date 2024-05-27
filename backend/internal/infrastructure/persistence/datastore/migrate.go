package datastore

import (
	"fmt"
	"myapp/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type MySQLMigrator struct {
	client *migrate.Migrate
}

func MustNewMySQLMigrator(migrateFilePath string) *MySQLMigrator {
	client, err := migrate.New(migrateFilePath, createDSN())
	if err != nil {
		panic(err)
	}
	return &MySQLMigrator{
		client: client,
	}
}

func createDSN() string {
	host := config.DBHostName
	port := config.DBPort
	dbname := config.DBName
	return fmt.Sprintf("mysql://root@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", host, port, dbname)
}
func (m *MySQLMigrator) Migrate() error {
	return m.client.Up()
}
