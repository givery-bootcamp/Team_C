package model_test

import (
	"myapp/internal/domain/model"
	"myapp/internal/pkg/test"
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewPost(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		post := model.NewPost(
		        "title",
		        "body",
		        model.User{
			        Name:     "test",
			        Password: "password",
		        },
		)

		expectedUser := test.DiffEq(&model.Post{
			Title: "title",
			Body:  "body",
			User: model.User{
				Name:     "test",
				Password: "password",
			},
		}, cmpopts.IgnoreFields(model.Post{}, "CreatedAt", "UpdatedAt"))

		if !expectedUser.Matches(post) {
			t.Errorf("got unexpected post: %+v", expectedUser.String())
		}
	})
}

func TestUpdatePost(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		post := model.Post{
			Title: "title",
			Body:  "body",
			User: model.User{
				Name:     "test",
				Password: "password",
			},
		}

		post.UpdatePost("new title", "new body")

		expectedUser := test.DiffEq(model.Post{
			Title: "new title",
			Body:  "new body",
			User: model.User{
				Name:     "test",
				Password: "password",
			},
		}, cmpopts.IgnoreFields(model.Post{}, "CreatedAt", "UpdatedAt"))

		if !expectedUser.Matches(post) {
			t.Errorf("got unexpected post: %+v", expectedUser.String())
		}
	})
}
