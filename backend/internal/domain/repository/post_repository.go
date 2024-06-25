//go:generate mockgen -source=post_repository.go -destination=repository_mock/post_repository_mock.go -package repository_mock
package repository

import (
	"context"

	"myapp/internal/domain/model"
)

type PostRepository interface {
	GetAll(ctx context.Context, limit, offset int) ([]*model.Post, error)
	GetByID(ctx context.Context, id int) (*model.Post, error)
	Create(ctx context.Context, post *model.Post) (*model.Post, error)
	Update(ctx context.Context, post *model.Post) (*model.Post, error)
	Delete(ctx context.Context, id int) error
}
