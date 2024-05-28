package usecase

import (
	"context"
	"myapp/internal/domain/model"
	"myapp/internal/domain/repository"
	"myapp/internal/exception"
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

func (u *UserUsecase) GetByID(ctx context.Context, id int) (*model.User, error) {
	user, err := u.r.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, exception.ServerError
	}

	return user, nil
}
