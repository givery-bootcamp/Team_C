package usecase

import (
	"context"
	"myapp/internal/domain/model"
	"myapp/internal/domain/repository"
)

type UserUsecase struct {
	r repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return UserUsecase{
		r: r,
	}
}

func (u *UserUsecase) Signin(ctx context.Context, param model.UserSigninParam) (*model.User, error) {
	return u.r.Signin(ctx, param)
}
