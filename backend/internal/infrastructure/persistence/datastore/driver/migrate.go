package driver

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type MySQLMigrator struct {
	client *migrate.Migrate
}

func MustNewMySQLMigrator(migrateFilePath string) *MySQLMigrator {
	client, err := migrate.New(migrateFilePath, createDSNForGoMigrate())
	if err != nil {
		panic(err)
	}
	return &MySQLMigrator{
		client: client,
	}
}

func (m *MySQLMigrator) Migrate() error {
	err := m.client.Up()

	if err != migrate.ErrNoChange {
		return err
	}
	return nil
}
