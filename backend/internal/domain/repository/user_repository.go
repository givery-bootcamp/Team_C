package repository

import (
	"context"
	"myapp/internal/domain/model"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int) (*model.User, error)
	GetBySigninParam(ctx context.Context, param model.UserSigninParam) (*model.User, error)
}
