package repository

import (
	"context"
	"myapp/internal/domain/model"
)

type PostRepository interface {
	GetAll(ctx context.Context) ([]*model.Post, error)
}
