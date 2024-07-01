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
					if r := recover(); r == nil {
						t.Errorf("The code did not panic as expected")
					}
				}()
			}
			_ = MustNewMySQLMigrator(tt.input)
		})
	}
}
