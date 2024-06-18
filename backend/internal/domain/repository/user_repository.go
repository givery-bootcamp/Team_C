package repository

import (
	"context"
	"myapp/internal/domain/model"
)

//go:generate mockgen -source=user_repository.go -destination=mock/user_repository_mock.go
type UserRepository interface {
	GetByID(ctx context.Context, id int) (*model.User, error)
	GetBySigninParam(ctx context.Context, param model.UserSigninParam) (*model.User, error)
}
