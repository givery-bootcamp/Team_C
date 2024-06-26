//go:generate mockgen -source=hello_world_repository.go -destination=repository_mock/hello_world_repository_mock.go -package repository_mock
package repository

import (
	"context"
	"myapp/internal/domain/model"
)

type HelloWorldRepository interface {
	Get(ctx context.Context, lang string) (*model.HelloWorld, error)
}
