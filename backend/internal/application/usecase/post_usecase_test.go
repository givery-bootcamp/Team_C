package usecase_test

import (
	"context"
	"errors"
	"myapp/internal/application/usecase"
	"myapp/internal/domain/model"
	mock_repository "myapp/internal/domain/repository/mock"
	"myapp/internal/pkg/test"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type testDependencies struct {
	r    *mock_repository.MockPostRepository
	u    usecase.PostUsecase
	ctrl *gomock.Controller
}

func newTestDependencies(t *testing.T) testDependencies {
	ctrl := gomock.NewController(t)
	r := mock_repository.NewMockPostRepository(ctrl)
	u := usecase.NewPostUsecase(r)

	return testDependencies{
		r:    r,
		u:    u,
		ctrl: ctrl,
	}
}

func TestPostUsecaseTest(t *testing.T) {
	t.Run("GetAll!!", func(t *testing.T) {
		tests := []struct {
			name          string
			limit         int
			offset        int
			ctx           context.Context
			posts         []*model.Post
			repositoryErr error
		}{
			{
				name:   "success",
				limit:  10,
				offset: 0,
				ctx:    context.Background(),
				posts: []*model.Post{
					model.NewPost("title1", "body1", *model.NewUser("user1", "password1")),
				},
				repositoryErr: nil,
			},
			{
				name:          "fail/Repositoryがエラーを返した時にエラーを返す",
				limit:         10,
				offset:        0,
				ctx:           context.Background(),
				posts:         nil,
				repositoryErr: errors.New("error"),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				deps := newTestDependencies(t)

				deps.r.EXPECT().GetAll(tt.ctx, tt.limit, tt.offset).Return(tt.posts, tt.repositoryErr)
				posts, err := deps.u.GetAll(tt.ctx, tt.limit, tt.offset)

				assert.Equal(t, tt.posts, posts)
				assert.Equal(t, tt.repositoryErr, err)
			})
		}
	})

	t.Run("GetByID", func(t *testing.T) {
		tests := []struct {
			name          string
			id            int
			ctx           context.Context
			post          *model.Post
			repositoryErr error
		}{
			{
				name:          "success",
				id:            1,
				ctx:           context.Background(),
				post:          model.NewPost("title1", "body1", *model.NewUser("user1", "password1")),
				repositoryErr: nil,
			},
			{
				name:          "fail/Repositoryがエラーを返した時にエラーを返す",
				id:            1,
				ctx:           context.Background(),
				post:          nil,
				repositoryErr: errors.New("error"),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				deps := newTestDependencies(t)

				deps.r.EXPECT().GetByID(tt.ctx, tt.id).Return(tt.post, tt.repositoryErr)
				post, err := deps.u.GetByID(tt.ctx, tt.id)

				assert.Equal(t, tt.post, post)
				assert.Equal(t, tt.repositoryErr, err)
			})
		}
	})

	t.Run("Create", func(t *testing.T) {
		tests := []struct {
			name          string
			title         string
			body          string
			userId        int
			ctx           context.Context
			inputtedPost  *model.Post
			returnedPost  *model.Post
			repositoryErr error
		}{
			{
				name:          "success",
				title:         "title1",
				body:          "body1",
				userId:        1,
				ctx:           context.Background(),
				inputtedPost:  model.NewPost("title1", "body1", model.User{ID: 1}),
				returnedPost:  model.NewPost("title1", "body1", *model.NewUser("user1", "password1")),
				repositoryErr: nil,
			},
			{
				name:          "fail/Repositoryがエラーを返した時にエラーを返す",
				title:         "title1",
				body:          "body1",
				userId:        1,
				ctx:           context.Background(),
				inputtedPost:  model.NewPost("title1", "body1", model.User{ID: 1}),
				returnedPost:  nil,
				repositoryErr: errors.New("error"),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				deps := newTestDependencies(t)

				postIgnoreFieldsOpt := cmp.Options{
					// ID, CreatedAt, UpdatedAtはUsecase内で生成されるため比較対象外にする
					cmpopts.IgnoreFields(model.Post{}, "ID", "CreatedAt", "UpdatedAt"),
				}
				deps.r.EXPECT().
					Create(tt.ctx, test.DiffEq(tt.inputtedPost, postIgnoreFieldsOpt)).
					Return(tt.returnedPost, tt.repositoryErr)
				post, err := deps.u.Create(tt.ctx, tt.title, tt.body, tt.userId)

				assert.Equal(t, tt.returnedPost, post)
				assert.Equal(t, tt.repositoryErr, err)
			})
		}
	})
}
