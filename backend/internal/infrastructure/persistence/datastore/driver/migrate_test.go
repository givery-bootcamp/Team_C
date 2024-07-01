package driver

import (
	"testing"
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
					err := recover()
					if err != "no scheme" {
						t.Errorf("got %v\nwant %v", err, "no scheme")
					}
				}()
			}
			_ = MustNewMySQLMigrator(tt.input)
		})
	}
}
