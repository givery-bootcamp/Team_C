package driver

import (
	"testing"
)

func TestMySQLMigrator_MustNewSQLMigrate(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr error
	}{
		{
			name:    "Successful migration",
			input:   "dummy",
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr != nil {
				_ = MustNewMySQLMigrator(tt.input)
			} else {
				defer func() {
					err := recover()
					if err != "no scheme" {
						t.Errorf("got %v\nwant %v", err, "no scheme")
					}
				}()
				_ = MustNewMySQLMigrator(tt.input)
			}
		})
	}
}
