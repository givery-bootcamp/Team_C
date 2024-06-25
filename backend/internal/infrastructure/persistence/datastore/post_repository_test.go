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

func TestPostRepository_GetAll(t *testing.T) {
	t.Run("failed/postsテーブルのクエリに失敗したとき期待したエラーを返す", func(t *testing.T) {
		deps, err := newTestPostRepositoryDependencies(t)
		require.NoError(t, err)

		deps.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts` LIMIT ? OFFSET ?")).
			WithArgs(10, 1).
			WillReturnError(gorm.ErrInvalidDB)

		posts, err := deps.repo.GetAll(context.Background(), 10, 1)

		AssertErrMsg(t, xerrors.Errorf("failed to SQL execution: %w", gorm.ErrInvalidDB), err)
		assert.Nil(t, posts)
	})

	t.Run("success", func(t *testing.T) {
		deps, err := newTestPostRepositoryDependencies(t)
		require.NoError(t, err)

		deps.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts` LIMIT ? OFFSET ?")).
			WithArgs(10, 1).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "title", "body", "user_id", "created_at", "updated_at", "deleted_at"}).
					AddRow(1, "タイトル1", "本文1", 11, time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local), nil).
					AddRow(2, "タイトル2", "本文2", 12, time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local), nil),
			)
		deps.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`id` IN (?,?)")).
			WithArgs(11, 12).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "password", "created_at", "updated_at", "deleted_at"}).
					AddRow(11, "ユーザー11", "password", time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local), nil).
					AddRow(12, "ユーザー12", "password", time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local), nil),
			)

		posts, err := deps.repo.GetAll(context.Background(), 10, 1)

		require.NoError(t, err)
		assert.Equal(t, []*model.Post{
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
		}, posts)
	})
}

func TestPostRepository_GetById(t *testing.T) {
	t.Run("failed/postsテーブルのクエリに失敗したとき期待したエラーを返す", func(t *testing.T) {
		deps, err := newTestPostRepositoryDependencies(t)
		require.NoError(t, err)

		deps.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts` WHERE id = ? ORDER BY `posts`.`id` LIMIT ?")).
			WithArgs(1, 1).
			WillReturnError(gorm.ErrInvalidDB)

		post, err := deps.repo.GetByID(context.Background(), 1)

		AssertErrMsg(t, xerrors.Errorf("failed to SQL execution: %w", gorm.ErrInvalidDB), err)
		assert.Nil(t, post)
	})

	t.Run("failed/指定したIDの投稿が存在しないとき期待したエラーを返す", func(t *testing.T) {
		deps, err := newTestPostRepositoryDependencies(t)
		require.NoError(t, err)

		deps.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts` WHERE id = ? ORDER BY `posts`.`id` LIMIT ?")).
			WithArgs(1, 1).
			WillReturnError(gorm.ErrRecordNotFound)

		post, err := deps.repo.GetByID(context.Background(), 1)

		AssertErrMsg(t, exception.RecordNotFoundError, err)
		assert.Nil(t, post)
	})

	t.Run("success", func(t *testing.T) {
		deps, err := newTestPostRepositoryDependencies(t)
		require.NoError(t, err)

		deps.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts` WHERE id = ? ORDER BY `posts`.`id` LIMIT ?")).
			WithArgs(1, 1).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "title", "body", "user_id", "created_at", "updated_at", "deleted_at"}).
					AddRow(1, "タイトル", "本文", 11, time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local), nil),
			)
		deps.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`id` = ?")).
			WithArgs(11).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "password", "created_at", "updated_at", "deleted_at"}).
					AddRow(11, "ユーザー11", "password", time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local), nil),
			)

		post, err := deps.repo.GetByID(context.Background(), 1)

		require.NoError(t, err)
		assert.Equal(t, &model.Post{
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
		}, post)
	})
}

func TestPostRepository_Create(t *testing.T) {
	t.Run("failed/postsテーブルのクエリに失敗したとき期待したエラーを返す", func(t *testing.T) {
		deps, err := newTestPostRepositoryDependencies(t)
		require.NoError(t, err)

		deps.mock.ExpectBegin()
		deps.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`name`,`password`,`created_at`,`updated_at`,`deleted_at`,`id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id`=`id`")).
			WithArgs("ユーザー", "password", sqlmock.AnyArg(), sqlmock.AnyArg(), nil, 11).
			WillReturnResult(sqlmock.NewResult(1, 1))

		deps.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `posts` (`title`,`body`,`user_id`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?)")).
			WithArgs("タイトル", "本文", 11, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnError(gorm.ErrInvalidDB)

		deps.mock.ExpectRollback()

		post, err := deps.repo.Create(context.Background(), &model.Post{Title: "タイトル", Body: "本文", User: model.User{ID: 11, Name: "ユーザー", Password: "password"}})

		AssertErrMsg(t, xerrors.Errorf("failed to SQL execution: %w", gorm.ErrInvalidDB), err)
		assert.Nil(t, post)
	})

	t.Run("failed/usersテーブルのSELECTクエリに失敗したとき期待したエラーを返す", func(t *testing.T) {
		deps, err := newTestPostRepositoryDependencies(t)
		require.NoError(t, err)

		deps.mock.ExpectBegin()
		deps.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`name`,`password`,`created_at`,`updated_at`,`deleted_at`,`id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id`=`id`")).
			WithArgs("ユーザー", "password", sqlmock.AnyArg(), sqlmock.AnyArg(), nil, 11).
			WillReturnResult(sqlmock.NewResult(1, 1))

		deps.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `posts` (`title`,`body`,`user_id`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?)")).
			WithArgs("タイトル", "本文", 11, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
		deps.mock.ExpectCommit()

		deps.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT ?")).
			WithArgs(11, 1).
			WillReturnError(gorm.ErrInvalidDB)

		post, err := deps.repo.Create(context.Background(), &model.Post{Title: "タイトル", Body: "本文", User: model.User{ID: 11, Name: "ユーザー", Password: "password"}})

		AssertErrMsg(t, xerrors.Errorf("failed to SQL execution: %w", gorm.ErrInvalidDB), err)
		assert.Nil(t, post)
	})

	t.Run("success", func(t *testing.T) {
		deps, err := newTestPostRepositoryDependencies(t)
		require.NoError(t, err)

		deps.mock.ExpectBegin()
		deps.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`name`,`password`,`created_at`,`updated_at`,`deleted_at`,`id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id`=`id`")).
			WithArgs("ユーザー", "password", sqlmock.AnyArg(), sqlmock.AnyArg(), nil, 11).
			WillReturnResult(sqlmock.NewResult(1, 1))

		deps.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `posts` (`title`,`body`,`user_id`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?)")).
			WithArgs("タイトル", "本文", 11, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
		deps.mock.ExpectCommit()

		deps.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT ?")).
			WithArgs(11, 1).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "password", "created_at", "updated_at", "deleted_at"}).
					AddRow(11, "ユーザー", "password", time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local), nil),
			)

		post, err := deps.repo.Create(context.Background(), &model.Post{Title: "タイトル", Body: "本文", User: model.User{ID: 11, Name: "ユーザー", Password: "password"}})
		assert.NoError(t, err)

		expectedPost := test.DiffEq(&model.Post{
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
		}, cmpopts.IgnoreFields(model.Post{}, "CreatedAt", "UpdatedAt"))

		if !expectedPost.Matches(post) {
			t.Errorf("got unexpected post: %v", expectedPost.String())
		}
	})
}

func TestPostRepository_Update(t *testing.T) {
	t.Run("failed/postsテーブルのクエリに失敗したとき期待したエラーを返す", func(t *testing.T) {
		deps, err := newTestPostRepositoryDependencies(t)
		require.NoError(t, err)

		deps.mock.ExpectBegin()
		deps.mock.ExpectExec(regexp.QuoteMeta("UPDATE `posts` SET `id`=?,`title`=?,`body`=?,`user_id`=?,`created_at`=?,`updated_at`=? WHERE id = ?")).
			WithArgs(1, "タイトル", "本文", 1, sqlmock.AnyArg(), sqlmock.AnyArg(), 1).
			WillReturnError(gorm.ErrInvalidDB)

		deps.mock.ExpectRollback()

		post, err := deps.repo.Update(context.Background(), &model.Post{
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
		})

		AssertErrMsg(t, xerrors.Errorf("failed to SQL execution: %w", gorm.ErrInvalidDB), err)
		assert.Nil(t, post)
	})

	t.Run("success", func(t *testing.T) {
		deps, err := newTestPostRepositoryDependencies(t)
		require.NoError(t, err)

		deps.mock.ExpectBegin()
		deps.mock.ExpectExec(regexp.QuoteMeta("UPDATE `posts` SET `id`=?,`title`=?,`body`=?,`user_id`=?,`created_at`=?,`updated_at`=? WHERE id = ?")).
			WithArgs(1, "タイトル", "本文", 1, sqlmock.AnyArg(), sqlmock.AnyArg(), 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		deps.mock.ExpectCommit()

		deps.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT ?")).
			WithArgs(1, 1).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "password", "created_at", "updated_at", "deleted_at"}).
					AddRow(1, "ユーザー", "password", time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local), nil),
			)

		post, err := deps.repo.Update(context.Background(), &model.Post{
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
		})
		assert.NoError(t, err)

		assert.Equal(t, &model.Post{
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
		}, post)
	})
}
