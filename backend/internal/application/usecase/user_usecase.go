package usecase

import (
	"context"
	"myapp/internal/domain/model"
	"myapp/internal/domain/repository"
	"myapp/internal/exception"
	"myapp/internal/pkg/hash"

	"golang.org/x/xerrors"
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
	user, err := u.r.GetByName(ctx, param.Name)
	if err != nil {
		return nil, err
	}

	err = hash.CompareHashPassword(user.Password, param.Password)
	if err != nil {
		return nil, xerrors.Errorf("failed to sign in: %w", exception.FailedToSigninError)
	}

	return user, nil
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

func (u *UserUsecase) Create(ctx context.Context, param model.CreateUserParam) (*model.User, error) {
	password, err := hash.GenerateHashPassword(param.Password)
	if err != nil {
		return nil, err
	}

	user := model.NewUser(param.Name, password)

	return u.r.Create(ctx, *user)
}
