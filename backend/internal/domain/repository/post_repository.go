//go:generate mockgen -source=post_repository.go -destination=mock/post_repository_mock.go -package mock
package repository

import (
	"context"

	"myapp/internal/domain/model"
)

type PostRepository interface {
	GetAll(ctx context.Context, limit, offset int) ([]*model.Post, error)
	GetByID(ctx context.Context, id int) (*model.Post, error)
	Create(ctx context.Context, post *model.Post) (*model.Post, error)
}
