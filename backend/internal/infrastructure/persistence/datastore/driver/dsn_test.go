package driver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDSNForGoMigrate(t *testing.T) {
	dsn := createDSNForGoMigrate()

	assert.Equal(t, "mysql://root:password@tcp(db:3306)/training?charset=utf8mb4&parseTime=True&loc=Local", dsn)
}

func TestCreateDSNForGorm(t *testing.T) {
	dsn := createDSNForGorm()

	assert.Equal(t, "root:password@tcp(db:3306)/training?charset=utf8mb4&parseTime=True&loc=Local", dsn)
}
