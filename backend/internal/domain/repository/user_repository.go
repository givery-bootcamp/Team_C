package repository

import (
	"context"
	"myapp/internal/domain/model"
)

type UserRepository interface {
	Signin(ctx context.Context, param model.UserSigninParam) (*model.User, error)
}
