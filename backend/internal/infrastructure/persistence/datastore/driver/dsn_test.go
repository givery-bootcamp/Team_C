package driver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDSNForGoMigrate(t *testing.T) {
	dsn := createDSNForGoMigrate()

	assert.Equal(t, "mysql://root:password@tcp(db:3306)/training?charset=utf8mb4&parseTime=True&loc=Local", dsn)
}
