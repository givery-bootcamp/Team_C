package driver

import (
	"context"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB interface {
	Transaction(ctx context.Context, f func(ctx context.Context) error) error
	GetDB(ctx context.Context) *gorm.DB
}

type db struct {
	conn *gorm.DB
}

func NewDB() *db {
	return &db{
		conn: initDB(),
	}
}

type txContextKey struct{}

var txKey = txContextKey{}

func (d *db) Transaction(ctx context.Context, f func(ctx context.Context) error) error {
	return d.conn.Transaction(func(tx *gorm.DB) error {
		ctx := context.WithValue(ctx, txKey, tx)
		return f(ctx)
	})
}

func (d *db) GetDB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(txKey).(*gorm.DB)
	if ok && tx != nil {
		return tx.WithContext(ctx)
	}

	return d.conn.WithContext(ctx)
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(createDSNForGorm()), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return db
}
