package datastore_test

import (
	"context"
	"myapp/internal/domain/model"
	"myapp/internal/domain/repository"
	"myapp/internal/infrastructure/persistence/datastore"
	"myapp/internal/infrastructure/persistence/datastore/driver/driver_mock"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

type testPostRepositoryDependencies struct {
	ctrl *gomock.Controller
	repo repository.PostRepository
}

func newTestPostRepositoryDependencies(t *testing.T, db *gorm.DB) (testPostRepositoryDependencies, error) {
	ctrl := gomock.NewController(t)
	driverMock := driver_mock.NewMockDB(ctrl)

	driverMock.EXPECT().GetDB(gomock.Any()).Return(db).AnyTimes()

	repo := datastore.NewPostRepository(driverMock)

	return testPostRepositoryDependencies{
		ctrl: ctrl,
		repo: repo,
	}, nil
}

func TestPostRepository(t *testing.T) {
	t.Run("GetAll", func(t *testing.T) {
		mockDB, mock, err := NewDbMock()
		require.NoError(t, err)

		type input struct {
			limit  int
			offset int
		}

		tests := []struct {
			name  string
			input input

			expectQuery         []*sqlmock.ExpectedQuery
			expectGormErr       error
			expectedFunctionErr error
			expectReturnedPosts []*model.Post
		}{
			{
				name:  "failed/postsテーブルのクエリに失敗したときエラーを返す",
				input: input{limit: 10, offset: 1},
				expectQuery: []*sqlmock.ExpectedQuery{
					mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts` LIMIT ? OFFSET ?")).
						WithArgs(10, 1).
						WillReturnError(gorm.ErrInvalidDB),
				},
				expectGormErr:       gorm.ErrInvalidDB,
				expectedFunctionErr: xerrors.Errorf("failed to SQL execution: %w", gorm.ErrInvalidDB),
				expectReturnedPosts: nil,
			},
			{
				name:  "failed/usersテーブルのクエリに失敗したときエラーを返す",
				input: input{limit: 10, offset: 1},
				expectQuery: []*sqlmock.ExpectedQuery{
					mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts` LIMIT ? OFFSET ?")).
						WithArgs(10, 1).
						WillReturnRows(sqlmock.NewRows([]string{"id", "title", "body", "user_id"}).AddRow(1, "タイトル", "本文", 1)),
					mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`id` = ?")).
						WithArgs(1).
						WillReturnError(gorm.ErrInvalidDB),
				},
				expectedFunctionErr: xerrors.Errorf("failed to SQL execution: %w", gorm.ErrInvalidDB),
				expectGormErr:       gorm.ErrInvalidDB,
				expectReturnedPosts: nil,
			},
			{
				name:  "success",
				input: input{limit: 10, offset: 1},
				expectQuery: []*sqlmock.ExpectedQuery{
					mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts` LIMIT ? OFFSET ?")).
						WithArgs(10, 1).
						WillReturnRows(
							sqlmock.NewRows([]string{"id", "title", "body", "user_id"}).
								AddRow(1, "タイトル1", "本文1", 11).
								AddRow(2, "タイトル2", "本文2", 12),
						),
					mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`id` IN (?,?)")).
						WithArgs(11, 12).
						WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(11, "ユーザー11").AddRow(12, "ユーザー12")),
				},
				expectedFunctionErr: nil,
				expectGormErr:       nil,
				expectReturnedPosts: []*model.Post{
					{
						ID:    1,
						Title: "タイトル1",
						Body:  "本文1",
						User: model.User{
							ID:   11,
							Name: "ユーザー11",
						},
					},
					{
						ID:    2,
						Title: "タイトル2",
						Body:  "本文2",
						User: model.User{
							ID:   12,
							Name: "ユーザー12",
						},
					},
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				deps, err := newTestPostRepositoryDependencies(t, mockDB)
				require.NoError(t, err)

				posts, err := deps.repo.GetAll(context.Background(), tt.input.limit, tt.input.offset)

				AssertErrMsg(t, tt.expectedFunctionErr, err)
				assert.Equal(t, tt.expectReturnedPosts, posts)
			})
		}
	})
}
