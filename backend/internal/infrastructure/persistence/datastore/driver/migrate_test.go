package driver

import (
	"errors"
	"myapp/internal/infrastructure/persistence/datastore/driver/migrate_mock"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"go.uber.org/mock/gomock"
)

func TestMySQLMigrator_MustNewSQLMigrate(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "Error no scheme",
			input:   "dummy",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic as expected")
					}
				}()
			}
			_ = MustNewMySQLMigrator(tt.input)
		})
	}
}

func TestMigrate(t *testing.T) {
	tests := []struct {
		name    string
		mockErr error
		wantErr bool
	}{
		{
			name:    "failed/ErrNoChange",
			mockErr: migrate.ErrNoChange,
			wantErr: false,
		},
		{
			name:    "failed/other error",
			mockErr: errors.New("error"),
			wantErr: true,
		},
		{
			name:    "Success",
			mockErr: nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			migrateClientMock := migrate_mock.NewMockMigrateClient(ctrl)
			migrateClientMock.EXPECT().Up().Return(tt.mockErr)

			m := MySQLMigrator{
				client: migrateClientMock,
			}
			err := m.Migrate()
			if (err != nil) != tt.wantErr {
				t.Errorf("wantErr: %v, got: %v", tt.wantErr, err)
			}
		})
	}
}
