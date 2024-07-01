package driver

import (
	"testing"
)

func TestInitDB_Error(t *testing.T) {
	t.Run(tt.name, func(t *testing.T) {
		if tt.wantErr {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("The code did not panic as expected")
				}
			}()
		}
		_ = initDB()
	})
}
