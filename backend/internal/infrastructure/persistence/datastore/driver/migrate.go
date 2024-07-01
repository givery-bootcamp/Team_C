//go:generate mockgen -source=migrate.go -destination=migrate_mock/migrate_mock.go -package migrate_mock
package driver

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"golang.org/x/xerrors"
)

type MySQLMigrator struct {
	client MigrateClient
}

type MigrateClient interface {
	Up() error
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
	if err != nil {
		if err == migrate.ErrNoChange {
			return nil
		}
		return xerrors.Errorf("failed to migrate: %w", err)
	}
	return nil
}
