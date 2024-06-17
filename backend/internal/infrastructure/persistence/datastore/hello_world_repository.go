package datastore

import (
	"context"
	"errors"
	"myapp/internal/domain/model"
	"myapp/internal/domain/repository"
	"myapp/internal/infrastructure/persistence/datastore/driver"
	"myapp/internal/infrastructure/persistence/datastore/entity"

	"gorm.io/gorm"
)

type HelloWorldRepository struct {
	db driver.DB
}

func NewHelloWorldRepository(db driver.DB) repository.HelloWorldRepository {
	return &HelloWorldRepository{
		db: db,
	}
}

func (r *HelloWorldRepository) Get(ctx context.Context, lang string) (*model.HelloWorld, error) {
	var obj entity.HelloWorld

	conn := r.db.GetDB(ctx)
	result := conn.Where("lang = ?", lang).First(&obj)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, NewSQLError(result.Error)
	}
	return obj.ToModel(), nil
}
