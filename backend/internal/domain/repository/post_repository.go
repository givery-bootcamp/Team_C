package repository

import (
	"context"
	"myapp/internal/domain/model"
)

type PostRepository interface {
	GetAll(ctx context.Context) ([]*model.Post, error)
	GetByID(ctx context.Context, id int) (*model.Post, error)
	Create(ctx context.Context, post *model.Post) (*model.Post, error)
}
