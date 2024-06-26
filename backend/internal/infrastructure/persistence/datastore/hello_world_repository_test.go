package datastore_test

import (
	"context"
	"myapp/internal/domain/repository"
	"myapp/internal/infrastructure/persistence/datastore"
	"myapp/internal/infrastructure/persistence/datastore/driver/driver_mock"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

type testHelloWorldRepositoryDependencies struct {
	ctrl *gomock.Controller
	mock sqlmock.Sqlmock
	repo repository.HelloWorldRepository
}

func newTestHelloWorldRepositoryDependencies(t *testing.T) (testHelloWorldRepositoryDependencies, error) {
	ctrl := gomock.NewController(t)
	driverMock := driver_mock.NewMockDB(ctrl)

	mockDB, mock, err := NewDbMock()
	if err != nil {
		return testHelloWorldRepositoryDependencies{}, err
	}

	driverMock.EXPECT().GetDB(gomock.Any()).Return(mockDB).AnyTimes()

	repo := datastore.NewHelloWorldRepository(driverMock)

	return testHelloWorldRepositoryDependencies{
		ctrl: ctrl,
		mock: mock,
		repo: repo,
	}, nil
}

func TestHelloWorldRepository_Get(t *testing.T) {
	t.Run("HelloWorldのレコードが見つからない場合nilを返す", func(t *testing.T) {
		deps, err := newTestHelloWorldRepositoryDependencies(t)
		require.NoError(t, err)

		deps.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `hello_worlds` WHERE lang = ? ORDER BY `hello_worlds`.`lang` LIMIT ?")).
			WithArgs("ja", 1).
			WillReturnError(gorm.ErrRecordNotFound)

		helloWorld, err := deps.repo.Get(context.Background(), "ja")

		assert.NoError(t, err)
		assert.Nil(t, helloWorld)
	})

	t.Run("hello_worldテーブルのクエリに失敗した時エラーを返す", func(t *testing.T) {
		deps, err := newTestHelloWorldRepositoryDependencies(t)
		require.NoError(t, err)

		deps.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `hello_worlds` WHERE lang = ? ORDER BY `hello_worlds`.`lang` LIMIT ?")).
			WithArgs("ja", 1).
			WillReturnError(gorm.ErrInvalidDB)

		helloWorld, err := deps.repo.Get(context.Background(), "ja")

		assert.Error(t, err)
		assert.Nil(t, helloWorld)
	})

	t.Run("success", func(t *testing.T) {
		deps, err := newTestHelloWorldRepositoryDependencies(t)
		require.NoError(t, err)

		deps.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `hello_worlds` WHERE lang = ? ORDER BY `hello_worlds`.`lang` LIMIT ?")).
			WithArgs("ja", 1).
			WillReturnRows(sqlmock.NewRows([]string{"lang", "message"}).AddRow("ja", "こんにちは"))

		helloWorld, err := deps.repo.Get(context.Background(), "ja")

		assert.NoError(t, err)
		assert.Equal(t, "こんにちは", helloWorld.Message)
	})
}
