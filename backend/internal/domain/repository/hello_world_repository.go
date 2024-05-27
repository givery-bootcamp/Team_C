package repository

import (
	"context"
	"myapp/internal/domain/model"
)

type HelloWorldRepository interface {
	Get(ctx context.Context, lang string) (*model.HelloWorld, error)
}
