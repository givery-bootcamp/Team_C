package usecase_test

import (
	"context"
	"errors"
	"myapp/internal/application/usecase"
	"myapp/internal/domain/model"
	"myapp/internal/domain/repository/repository_mock"
	"myapp/internal/exception"
	"myapp/internal/pkg/test"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type testDependencies struct {
	r    *repository_mock.MockPostRepository
	u    usecase.PostUsecase
	ctrl *gomock.Controller
}

func newTestDependencies(t *testing.T) testDependencies {
	ctrl := gomock.NewController(t)
	r := repository_mock.NewMockPostRepository(ctrl)
	u := usecase.NewPostUsecase(r)

	return testDependencies{
		r:    r,
		u:    u,
		ctrl: ctrl,
	}
}

func TestPostUsecaseTest(t *testing.T) {
	t.Run("GetAll", func(t *testing.T) {
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

	t.Run("Update", func(t *testing.T) {
		t.Run("r.GetByIDがエラーを返した時にエラーを返す", func(t *testing.T) {
			deps := newTestDependencies(t)

			deps.r.EXPECT().GetByID(context.Background(), 1).Return(nil, errors.New("error"))
			post, err := deps.u.Update(context.Background(), "title1", "body1", 1, 1)

			assert.Nil(t, post)
			assert.NotNil(t, err)
		})

		t.Run("UpdateするPostが自分のものではない場合にRecordNotFoundErrorを返す", func(t *testing.T) {
			deps := newTestDependencies(t)

			deps.r.EXPECT().GetByID(context.Background(), 1).Return(model.NewPost("title1", "body1", model.User{ID: 2}), nil)
			post, err := deps.u.Update(context.Background(), "title1", "body1", 1, 1)

			assert.Nil(t, post)
			assert.Equal(t, err, exception.RecordNotFoundError)
		})

		t.Run("r.Updateがエラーを返した時にエラーを返す", func(t *testing.T) {
			deps := newTestDependencies(t)

			deps.r.EXPECT().GetByID(context.Background(), 1).Return(model.NewPost("title1", "body1", model.User{ID: 1}), nil)
			deps.r.EXPECT().Update(
				context.Background(),
				test.DiffEq(model.NewPost("title1", "body1", model.User{ID: 1}),
					cmpopts.IgnoreFields(model.Post{}, "ID", "CreatedAt", "UpdatedAt"),
				)).
				Return(nil, errors.New("error"))
			post, err := deps.u.Update(context.Background(), "title1", "body1", 1, 1)

			assert.Nil(t, post)
			assert.NotNil(t, err)
		})

		t.Run("success", func(t *testing.T) {
			deps := newTestDependencies(t)

			deps.r.EXPECT().GetByID(context.Background(), 1).Return(model.NewPost("title1", "body1", model.User{ID: 1}), nil)
			deps.r.EXPECT().Update(
				context.Background(),
				test.DiffEq(model.NewPost("title1", "body1", model.User{ID: 1}),
					cmpopts.IgnoreFields(model.Post{}, "ID", "CreatedAt", "UpdatedAt"),
				)).
				Return(model.NewPost("title1", "body1", *model.NewUser("user1", "password1")), nil)
			post, err := deps.u.Update(context.Background(), "title1", "body1", 1, 1)

			assert.NotNil(t, post)
			assert.Nil(t, err)
		})
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("r.GetByIDがエラーを返した時にエラーを返す", func(t *testing.T) {
			deps := newTestDependencies(t)

			deps.r.EXPECT().GetByID(context.Background(), 1).Return(nil, errors.New("error"))
			err := deps.u.Delete(context.Background(), 1, 1)

			assert.NotNil(t, err)
		})

		t.Run("DeleteするPostが自分のものではない場合にRecordNotFoundErrorを返す", func(t *testing.T) {
			deps := newTestDependencies(t)

			deps.r.EXPECT().GetByID(context.Background(), 1).Return(model.NewPost("title1", "body1", model.User{ID: 2}), nil)
			err := deps.u.Delete(context.Background(), 1, 1)

			assert.Equal(t, err, exception.RecordNotFoundError)
		})

		t.Run("r.Deleteがエラーを返した時にエラーを返す", func(t *testing.T) {
			deps := newTestDependencies(t)

			deps.r.EXPECT().GetByID(context.Background(), 1).Return(model.NewPost("title1", "body1", model.User{ID: 1}), nil)
			deps.r.EXPECT().Delete(context.Background(), 1).Return(errors.New("error"))
			err := deps.u.Delete(context.Background(), 1, 1)

			assert.NotNil(t, err)
		})
	})
}
