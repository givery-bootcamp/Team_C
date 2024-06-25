package model_test

import (
	"myapp/internal/domain/model"
	"myapp/internal/pkg/test"
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewUser(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		user := model.NewUser("test", "password")

		expectedUser := test.DiffEq(user, cmpopts.IgnoreFields(*user, "CreatedAt", "UpdatedAt"))
		if !expectedUser.Matches(user) {
			t.Errorf("got unexpected user: %+v", user)
		}
	})
}
