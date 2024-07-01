package datastore_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	mockDB, err := gorm.Open(
		mysql.New(mysql.Config{
			DriverName:                "mysql",
			Conn:                      db,
			SkipInitializeWithVersion: true,
		}),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)

	return mockDB, mock, err
}
func AssertErrMsg(t *testing.T, expected error, err error) {
	if expected == nil {
		assert.NoError(t, err)
		return
	}

	require.Error(t, err)
	assert.Equal(t, expected.Error(), err.Error())
}
