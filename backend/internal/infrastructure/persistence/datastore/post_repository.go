package datastore

import (
	"context"
	"myapp/internal/domain/model"
	"myapp/internal/domain/repository"
	"myapp/internal/infrastructure/persistence/datastore/driver"
	"myapp/internal/infrastructure/persistence/datastore/entity"
)

type PostRepository struct {
	db driver.DB
}

func NewPostRepository(db driver.DB) repository.PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (r *PostRepository) GetAll(ctx context.Context) ([]*model.Post, error) {
	posts := []*entity.Post{}

	conn := r.db.GetDB(ctx)
	result := conn.Preload("User").Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return entity.ToPostModelListFromEntity(posts), nil
}
