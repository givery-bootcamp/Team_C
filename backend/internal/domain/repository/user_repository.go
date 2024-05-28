package repository

import (
	"context"
	"myapp/internal/domain/model"
)

type UserRepository interface {
	GetBySigninParam(ctx context.Context, param model.UserSigninParam) (*model.User, error)
}
