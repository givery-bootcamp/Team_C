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

func (r *PostRepository) GetAll(ctx context.Context, limit, offset int) ([]*model.Post, error) {
	posts := []*entity.Post{}

	conn := r.db.GetDB(ctx)
	if err := conn.Preload("User").Limit(limit).Offset(offset).Find(&posts).Error; err != nil {
		return nil, err
	}
	return entity.ToPostModelListFromEntity(posts), nil
}

func (r *PostRepository) GetByID(ctx context.Context, id int) (*model.Post, error) {
	var p entity.Post

	conn := r.db.GetDB(ctx)
	if err := conn.Preload("User").Where("id = ?", id).First(&p).Error; err != nil {
		return nil, err
	}

	return p.ToModel(), nil
}
