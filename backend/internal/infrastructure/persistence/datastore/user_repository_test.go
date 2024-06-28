package datastore_test

import (
	"context"
	"errors"
	"myapp/internal/domain/model"
	"myapp/internal/domain/repository"
	"myapp/internal/infrastructure/persistence/datastore"
	"myapp/internal/infrastructure/persistence/datastore/driver/driver_mock"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

type testUserRepositoryDependencies struct {
	ctrl *gomock.Controller
	mock sqlmock.Sqlmock
	repo repository.UserRepository
}

func newTestUserRepositoryDependencies(t *testing.T) (testUserRepositoryDependencies, error) {
	ctrl := gomock.NewController(t)
	driverMock := driver_mock.NewMockDB(ctrl)

	mockDB, mock, err := NewDbMock()
	if err != nil {
		return testUserRepositoryDependencies{}, err
	}

	driverMock.EXPECT().GetDB(gomock.Any()).Return(mockDB).AnyTimes()

	repo := datastore.NewUserRepository(driverMock)

	return testUserRepositoryDependencies{
		ctrl: ctrl,
		mock: mock,
		repo: repo,
	}, nil
}

func TestUserRepository_GetByName(t *testing.T) {
	t.Run("failed/Userのレコードが見つからない場合エラーを返す", func(t *testing.T) {
		deps, err := newTestUserRepositoryDependencies(t)
		require.NoError(t, err)

		deps.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE name = ? ORDER BY `users`.`id` LIMIT ?")).
			WithArgs("user", 1).
			WillReturnError(gorm.ErrRecordNotFound)

		_, err = deps.repo.GetByName(context.Background(), "user")

		assert.Equal(t, "failed to get sign in params: レコードが見つかりませんでした", err.Error())
	})

	t.Run("failed/usersテーブルのクエリに失敗した時エラーを返す", func(t *testing.T) {
		deps, err := newTestUserRepositoryDependencies(t)
		require.NoError(t, err)

		deps.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE name = ? ORDER BY `users`.`id` LIMIT ?")).
			WithArgs("user", 1).
			WillReturnError(gorm.ErrInvalidDB)

		_, err = deps.repo.GetByName(context.Background(), "user")

		assert.Equal(t, "failed to SQL execution: invalid db", err.Error())
	})

	t.Run("success", func(t *testing.T) {
		deps, err := newTestUserRepositoryDependencies(t)
		if err != nil {
			t.Fatal(err)
		}
		defer deps.ctrl.Finish()

		deps.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE name = ? ORDER BY `users`.`id` LIMIT ?")).
			WithArgs("user", 1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "password", "created_at", "updated_at"}).
				AddRow(1, "user", "password", time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			)

		user, err := deps.repo.GetByName(context.Background(), "user")

		assert.NoError(t, err)
		assert.Equal(t, &model.User{
			ID:        1,
			Name:      "user",
			Password:  "password",
			CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		}, user)
	})
}

func TestUserRepository_GetById(t *testing.T) {
	t.Run("failed/usersテーブルのクエリに失敗した時エラーを返す", func(t *testing.T) {
		deps, err := newTestUserRepositoryDependencies(t)
		require.NoError(t, err)

		deps.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT ?")).
			WithArgs(1, 1).
			WillReturnError(gorm.ErrInvalidDB)

		_, err = deps.repo.GetByID(context.Background(), 1)

		assert.Equal(t, "failed to SQL execution: invalid db", err.Error())
	})

	t.Run("success/Userのレコードが見つからない場合nilを返す", func(t *testing.T) {
		deps, err := newTestUserRepositoryDependencies(t)
		require.NoError(t, err)

		deps.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT ?")).
			WithArgs(1, 1).
			WillReturnError(gorm.ErrRecordNotFound)

		user, err := deps.repo.GetByID(context.Background(), 1)

		assert.NoError(t, err)
		assert.Nil(t, user)
	})

	t.Run("success/Userのレコードが見つかった場合それを返す", func(t *testing.T) {
		deps, err := newTestUserRepositoryDependencies(t)
		require.NoError(t, err)

		deps.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT ?")).
			WithArgs(1, 1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "password", "created_at", "updated_at"}).
				AddRow(1, "user", "password", time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			)

		user, err := deps.repo.GetByID(context.Background(), 1)

		assert.NoError(t, err)
		assert.Equal(t, &model.User{
			ID:        1,
			Name:      "user",
			Password:  "password",
			CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		}, user)
	})
}

func TestUserRepository_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		deps, err := newTestUserRepositoryDependencies(t)
		require.NoError(t, err)
		defer deps.ctrl.Finish()

		user := model.User{
			Name:     "new_user",
			Password: "password",
		}

		deps.mock.ExpectBegin()

		deps.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`name`,`password`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?)")).
			WithArgs(user.Name, user.Password, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))

		deps.mock.ExpectCommit()

		result, err := deps.repo.Create(context.Background(), user)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, user.Name, result.Name)
		assert.Equal(t, user.Password, result.Password)
		assert.NotZero(t, result.CreatedAt)
		assert.NotZero(t, result.UpdatedAt)
	})

	t.Run("failed/duplicate entry", func(t *testing.T) {
		deps, err := newTestUserRepositoryDependencies(t)
		require.NoError(t, err)
		defer deps.ctrl.Finish()

		user := model.User{
			Name:     "duplicate_user",
			Password: "password",
		}

		mySQLError := &mysql.MySQLError{
			Number:  1062,
			Message: "Duplicate entry",
		}

		deps.mock.ExpectBegin()

		deps.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`name`,`password`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?)")).
			WithArgs(user.Name, user.Password, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnError(mySQLError)

		deps.mock.ExpectRollback()

		_, err = deps.repo.Create(context.Background(), user)

		assert.EqualError(t, err, "failed to create user: すでに使われているユーザー名です")
	})

	t.Run("failed/sql execution error", func(t *testing.T) {
		deps, err := newTestUserRepositoryDependencies(t)
		require.NoError(t, err)
		defer deps.ctrl.Finish()

		user := model.User{
			Name:     "new_user",
			Password: "password",
		}

		sqlError := errors.New("some SQL error")

		deps.mock.ExpectBegin()

		deps.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`name`,`password`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?)")).
			WithArgs(user.Name, user.Password, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnError(sqlError)

		deps.mock.ExpectRollback()

		_, err = deps.repo.Create(context.Background(), user)

		assert.EqualError(t, err, "failed to SQL execution: some SQL error")
	})
}
