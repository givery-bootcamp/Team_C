//go:build !unit_test
// +build !unit_test

package e2e_test

import (
	"context"
	"testing"

	"github.com/k1LoW/runn"
)

func Test_runn(t *testing.T) {
	t.Run("E2E Test", func(t *testing.T) {
		opts := []runn.Option{
			runn.T(t),
			runn.Runner("req", "http://test-backend:9000"),
		}

		o, err := runn.Load("./scenario/*.yml", opts...)
		if err != nil {
			t.Fatal(err)
		}

		ctx := context.Background()
		if err := o.RunN(ctx); err != nil {
			t.Fatal(err)
		}
	})
}
