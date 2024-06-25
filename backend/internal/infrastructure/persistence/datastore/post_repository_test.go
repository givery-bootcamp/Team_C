package datastore_test

import (
	"context"
	"myapp/internal/domain/model"
	"myapp/internal/domain/repository"
	"myapp/internal/exception"
	"myapp/internal/infrastructure/persistence/datastore"
	"myapp/internal/infrastructure/persistence/datastore/driver/driver_mock"
	"myapp/internal/pkg/test"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

type testPostRepositoryDependencies struct {
	ctrl *gomock.Controller
	mock sqlmock.Sqlmock
	repo repository.PostRepository
}

func newTestPostRepositoryDependencies(t *testing.T) (testPostRepositoryDependencies, error) {
	ctrl := gomock.NewController(t)
	driverMock := driver_mock.NewMockDB(ctrl)

	mockDB, mock, err := NewDbMock()
	if err != nil {
		return testPostRepositoryDependencies{}, err
	}

	driverMock.EXPECT().GetDB(gomock.Any()).Return(mockDB).AnyTimes()

	repo := datastore.NewPostRepository(driverMock)

	return testPostRepositoryDependencies{
		ctrl: ctrl,
		mock: mock,
		repo: repo,
	}, nil
}

func TestPostRepository(t *testing.T) {
	t.Run("GetAll", func(t *testing.T) {
		type input struct {
			limit  int
			offset int
		}

		tests := []struct {
			name  string
			input input

			buildMockQueryFn    func(sqlmock.Sqlmock)
			expectQuery         []*sqlmock.ExpectedQuery
			expectedFunctionErr error
			expectReturnedPosts []*model.Post
		}{
			{
				name:  "failed/postsテーブルのクエリに失敗したとき期待したエラーを返す",
				input: input{limit: 10, offset: 1},
				buildMockQueryFn: func(mock sqlmock.Sqlmock) {
					mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts` LIMIT ? OFFSET ?")).
						WithArgs(10, 1).
						WillReturnError(gorm.ErrInvalidDB)
				},
				expectedFunctionErr: xerrors.Errorf("failed to SQL execution: %w", gorm.ErrInvalidDB),
				expectReturnedPosts: nil,
			},
			{
				name:  "failed/usersテーブルのクエリに失敗したとき期待したエラーを返す",
				input: input{limit: 10, offset: 1},
				buildMockQueryFn: func(mock sqlmock.Sqlmock) {
					mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts` LIMIT ? OFFSET ?")).
						WithArgs(10, 1).
						WillReturnRows(sqlmock.NewRows([]string{"id", "title", "body", "user_id", "created_at", "updated_at", "deleted_at"}).
							AddRow(1, "タイトル", "本文", 1, time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local), nil),
						)
					mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`id` = ?")).
						WithArgs(1).
						WillReturnError(gorm.ErrInvalidDB)
				},
				expectedFunctionErr: xerrors.Errorf("failed to SQL execution: %w", gorm.ErrInvalidDB),
				expectReturnedPosts: nil,
			},
			{
				name:  "success",
				input: input{limit: 10, offset: 1},
				buildMockQueryFn: func(mock sqlmock.Sqlmock) {
					mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts` LIMIT ? OFFSET ?")).
						WithArgs(10, 1).
						WillReturnRows(
							sqlmock.NewRows([]string{"id", "title", "body", "user_id", "created_at", "updated_at", "deleted_at"}).
								AddRow(1, "タイトル1", "本文1", 11, time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local), nil).
								AddRow(2, "タイトル2", "本文2", 12, time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local), nil),
						)
					mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`id` IN (?,?)")).
						WithArgs(11, 12).
						WillReturnRows(
							sqlmock.NewRows([]string{"id", "name", "password", "created_at", "updated_at", "deleted_at"}).
								AddRow(11, "ユーザー11", "password", time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local), nil).
								AddRow(12, "ユーザー12", "password", time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local), nil),
						)
				},
				expectedFunctionErr: nil,
				expectReturnedPosts: []*model.Post{
					{
						ID:    1,
						Title: "タイトル1",
						Body:  "本文1",
						User: model.User{
							ID:        11,
							Name:      "ユーザー11",
							Password:  "password",
							CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
							UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
						},
						CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
						UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
					},
					{
						ID:    2,
						Title: "タイトル2",
						Body:  "本文2",
						User: model.User{
							ID:        12,
							Name:      "ユーザー12",
							Password:  "password",
							CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
							UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
						},
						CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
						UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
					},
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				deps, err := newTestPostRepositoryDependencies(t)
				require.NoError(t, err)

				tt.buildMockQueryFn(deps.mock)

				posts, err := deps.repo.GetAll(context.Background(), tt.input.limit, tt.input.offset)

				AssertErrMsg(t, tt.expectedFunctionErr, err)
				assert.Equal(t, tt.expectReturnedPosts, posts)
			})
		}
	})

	t.Run("GetByID", func(t *testing.T) {
		type input struct {
			id int
		}

		tests := []struct {
			name  string
			input input

			buildMockQueryFn    func(sqlmock.Sqlmock)
			expectQuery         []*sqlmock.ExpectedQuery
			expectedFunctionErr error
			expectReturnedPost  *model.Post
		}{
			{
				name:  "failed/postsテーブルのクエリに失敗したとき期待したエラーを返す",
				input: input{id: 1},
				buildMockQueryFn: func(mock sqlmock.Sqlmock) {
					mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts` WHERE id = ? ORDER BY `posts`.`id` LIMIT ?")).
						WithArgs(1, 1).
						WillReturnError(gorm.ErrInvalidDB)
				},
				expectedFunctionErr: xerrors.Errorf("failed to SQL execution: %w", gorm.ErrInvalidDB),
				expectReturnedPost:  nil,
			},
			{
				name:  "failed/usersテーブルのクエリに失敗したとき期待したエラーを返す",
				input: input{id: 1},
				buildMockQueryFn: func(mock sqlmock.Sqlmock) {
					mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts` WHERE id = ? ORDER BY `posts`.`id` LIMIT ?")).
						WithArgs(1, 1).
						WillReturnRows(sqlmock.NewRows([]string{"id", "title", "body", "user_id", "created_at", "updated_at", "deleted_at"}).
							AddRow(1, "タイトル", "本文", 1, time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local), nil),
						)
					mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`id` = ?")).
						WithArgs(1).
						WillReturnError(gorm.ErrInvalidDB)
				},
				expectedFunctionErr: xerrors.Errorf("failed to SQL execution: %w", gorm.ErrInvalidDB),
				expectReturnedPost:  nil,
			},
			{
				name:  "failed/指定したIDの投稿が存在しないとき期待したエラーを返す",
				input: input{id: 1},
				buildMockQueryFn: func(mock sqlmock.Sqlmock) {
					mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts` WHERE id = ? ORDER BY `posts`.`id` LIMIT ?")).
						WithArgs(1, 1).
						WillReturnError(gorm.ErrRecordNotFound)
				},
				expectedFunctionErr: exception.RecordNotFoundError,
				expectReturnedPost:  nil,
			},
			{
				name:  "success",
				input: input{id: 1},
				buildMockQueryFn: func(mock sqlmock.Sqlmock) {
					mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts` WHERE id = ? ORDER BY `posts`.`id` LIMIT ?")).
						WithArgs(1, 1).
						WillReturnRows(
							sqlmock.NewRows([]string{"id", "title", "body", "user_id", "created_at", "updated_at", "deleted_at"}).
								AddRow(1, "タイトル", "本文", 11, time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local), nil),
						)
					mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`id` = ?")).
						WithArgs(11).
						WillReturnRows(
							sqlmock.NewRows([]string{"id", "name", "password", "created_at", "updated_at", "deleted_at"}).
								AddRow(11, "ユーザー11", "password", time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local), nil),
						)
				},
				expectedFunctionErr: nil,
				expectReturnedPost: &model.Post{
					ID:    1,
					Title: "タイトル",
					Body:  "本文",
					User: model.User{
						ID:        11,
						Name:      "ユーザー11",
						Password:  "password",
						CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
						UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
					},
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
					UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				deps, err := newTestPostRepositoryDependencies(t)
				require.NoError(t, err)

				tt.buildMockQueryFn(deps.mock)

				post, err := deps.repo.GetByID(context.Background(), tt.input.id)

				AssertErrMsg(t, tt.expectedFunctionErr, err)
				assert.Equal(t, tt.expectReturnedPost, post)
			})
		}
	})

	t.Run("Create", func(t *testing.T) {
		type input struct {
			post *model.Post
		}

		tests := []struct {
			name  string
			input input

			buildMockQueryFn    func(sqlmock.Sqlmock)
			expectQuery         []*sqlmock.ExpectedQuery
			expectedFunctionErr error
			expectReturnedPost  *model.Post
		}{
			{
				name:  "failed/usersテーブルのクエリに失敗したとき期待したエラーを返す",
				input: input{post: &model.Post{Title: "タイトル", Body: "本文", User: model.User{ID: 11, Name: "ユーザー", Password: "password"}}},
				buildMockQueryFn: func(mock sqlmock.Sqlmock) {
					mock.ExpectBegin()
					mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`name`,`password`,`created_at`,`updated_at`,`deleted_at`,`id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id`=`id`")).
						WithArgs("ユーザー", "password", sqlmock.AnyArg(), sqlmock.AnyArg(), nil, 11).
						WillReturnError(gorm.ErrInvalidDB)

					mock.ExpectRollback()
				},
				expectedFunctionErr: xerrors.Errorf("failed to SQL execution: %w", gorm.ErrInvalidDB),
				expectReturnedPost:  nil,
			},
			{
				name:  "failed/postsテーブルのクエリに失敗したとき期待したエラーを返す",
				input: input{post: &model.Post{Title: "タイトル", Body: "本文", User: model.User{ID: 11, Name: "ユーザー", Password: "password"}}},
				buildMockQueryFn: func(mock sqlmock.Sqlmock) {
					mock.ExpectBegin()
					mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`name`,`password`,`created_at`,`updated_at`,`deleted_at`,`id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id`=`id`")).
						WithArgs("ユーザー", "password", sqlmock.AnyArg(), sqlmock.AnyArg(), nil, 11).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `posts` (`title`,`body`,`user_id`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?)")).
						WithArgs("タイトル", "本文", 11, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
						WillReturnError(gorm.ErrInvalidDB)

					mock.ExpectRollback()
				},
				expectedFunctionErr: xerrors.Errorf("failed to SQL execution: %w", gorm.ErrInvalidDB),
				expectReturnedPost:  nil,
			},
			{
				name:  "failed/usersテーブルのSELECTに失敗したとき期待したエラーを返す",
				input: input{post: &model.Post{Title: "タイトル", Body: "本文", User: model.User{ID: 11, Name: "ユーザー", Password: "password"}}},
				buildMockQueryFn: func(mock sqlmock.Sqlmock) {
					mock.ExpectBegin()
					mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`name`,`password`,`created_at`,`updated_at`,`deleted_at`,`id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id`=`id`")).
						WithArgs("ユーザー", "password", sqlmock.AnyArg(), sqlmock.AnyArg(), nil, 11).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `posts` (`title`,`body`,`user_id`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?)")).
						WithArgs("タイトル", "本文", 11, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
						WillReturnResult(sqlmock.NewResult(1, 1))
					mock.ExpectCommit()

					mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT ?")).
						WithArgs(11, 1).
						WillReturnError(gorm.ErrInvalidDB)
				},
				expectedFunctionErr: xerrors.Errorf("failed to SQL execution: %w", gorm.ErrInvalidDB),
				expectReturnedPost:  nil,
			},
			{
				name:  "success",
				input: input{post: &model.Post{Title: "タイトル", Body: "本文", User: model.User{ID: 11, Name: "ユーザー", Password: "password"}}},
				buildMockQueryFn: func(mock sqlmock.Sqlmock) {
					mock.ExpectBegin()
					mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`name`,`password`,`created_at`,`updated_at`,`deleted_at`,`id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id`=`id`")).
						WithArgs("ユーザー", "password", sqlmock.AnyArg(), sqlmock.AnyArg(), nil, 11).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `posts` (`title`,`body`,`user_id`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?)")).
						WithArgs("タイトル", "本文", 11, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
						WillReturnResult(sqlmock.NewResult(1, 1))
					mock.ExpectCommit()

					mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT ?")).
						WithArgs(11, 1).
						WillReturnRows(
							sqlmock.NewRows([]string{"id", "name", "password", "created_at", "updated_at", "deleted_at"}).
								AddRow(11, "ユーザー", "password", time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local), nil),
						)
				},
				expectedFunctionErr: nil,
				expectReturnedPost: &model.Post{
					ID:    1,
					Title: "タイトル",
					Body:  "本文",
					User: model.User{
						ID:        11,
						Name:      "ユーザー",
						Password:  "password",
						CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
						UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
					},
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
					UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				deps, err := newTestPostRepositoryDependencies(t)
				require.NoError(t, err)

				tt.buildMockQueryFn(deps.mock)

				post, err := deps.repo.Create(context.Background(), tt.input.post)

				AssertErrMsg(t, tt.expectedFunctionErr, err)

				// CreatedAt, UpdatedAtはGorm内部で更新されるため比較対象外にする
				expectReturnedPost := test.DiffEq(tt.expectReturnedPost, cmpopts.IgnoreFields(model.Post{}, "CreatedAt", "UpdatedAt"))

				if !expectReturnedPost.Matches(post) {
					t.Errorf("expected: %v, but got: %v", tt.expectReturnedPost, post)
				}
			})
		}
	})

	t.Run("Update", func(t *testing.T) {
		type input struct {
			post *model.Post
		}

		tests := []struct {
			name  string
			input input

			buildMockQueryFn    func(sqlmock.Sqlmock)
			expectQuery         []*sqlmock.ExpectedQuery
			expectedFunctionErr error
			expectReturnedPost  *model.Post
		}{
			{
				name: "failed/postsテーブルのクエリに失敗したとき期待したエラーを返す",
				input: input{post: &model.Post{
					ID:    1,
					Title: "タイトル",
					Body:  "本文",
					User: model.User{
						ID:        1,
						Name:      "ユーザー",
						Password:  "password",
						CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
						UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
					},
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
					UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
				}},
				buildMockQueryFn: func(mock sqlmock.Sqlmock) {
					mock.ExpectBegin()
					mock.ExpectExec(regexp.QuoteMeta("UPDATE `posts` SET `id`=?,`title`=?,`body`=?,`user_id`=?,`created_at`=?,`updated_at`=? WHERE id = ?")).
						WithArgs(1, "タイトル", "本文", 1, sqlmock.AnyArg(), sqlmock.AnyArg(), 1).
						WillReturnError(gorm.ErrInvalidDB)

					mock.ExpectRollback()
				},
				expectedFunctionErr: xerrors.Errorf("failed to SQL execution: %w", gorm.ErrInvalidDB),
				expectReturnedPost:  nil,
			},
			{
				name: "success",
				input: input{post: &model.Post{
					ID:    1,
					Title: "タイトル",
					Body:  "本文",
					User: model.User{
						ID:        1,
						Name:      "ユーザー",
						Password:  "password",
						CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
						UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
					},
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
					UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
				}},
				buildMockQueryFn: func(mock sqlmock.Sqlmock) {
					mock.ExpectBegin()
					mock.ExpectExec(regexp.QuoteMeta("UPDATE `posts` SET `id`=?,`title`=?,`body`=?,`user_id`=?,`created_at`=?,`updated_at`=? WHERE id = ?")).
						WithArgs(1, "タイトル", "本文", 1, sqlmock.AnyArg(), sqlmock.AnyArg(), 1).
						WillReturnResult(sqlmock.NewResult(1, 1))
					mock.ExpectCommit()

					mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT ?")).
						WithArgs(1, 1).
						WillReturnRows(
							sqlmock.NewRows([]string{"id", "name", "password", "created_at", "updated_at", "deleted_at"}).
								AddRow(1, "ユーザー", "password", time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local), nil),
						)
				},
				expectedFunctionErr: nil,
				expectReturnedPost: &model.Post{
					ID:    1,
					Title: "タイトル",
					Body:  "本文",
					User: model.User{
						ID:        1,
						Name:      "ユーザー",
						Password:  "password",
						CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
						UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
					},
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
					UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				deps, err := newTestPostRepositoryDependencies(t)
				require.NoError(t, err)

				tt.buildMockQueryFn(deps.mock)

				post, err := deps.repo.Update(context.Background(), tt.input.post)

				AssertErrMsg(t, tt.expectedFunctionErr, err)
				assert.Equal(t, tt.expectReturnedPost, post)
			})
		}
	})
}
