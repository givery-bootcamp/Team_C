package datastore

import (
	"context"
	"errors"

	"myapp/internal/domain/model"
	"myapp/internal/domain/repository"
	"myapp/internal/exception"
	"myapp/internal/infrastructure/persistence/datastore/driver"
	"myapp/internal/infrastructure/persistence/datastore/entity"

	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

type UserRepository struct {
	db driver.DB
}

func NewUserRepository(db driver.DB) repository.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetBySigninParam(ctx context.Context, param model.UserSigninParam) (*model.User, error) {
	var user entity.User
	conn := r.db.GetDB(ctx)
	if err := conn.
		Where("name = ?", param.Name).
		Where("password = ?", param.Password).
		First(&user).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, xerrors.Errorf("failed to get sign in params: %w", exception.FailedToSigninError)
		}
		return nil, xerrors.Errorf("failed to SQL execution: %w", err)
	}

	return user.ToModel(), nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*model.User, error) {
	var user entity.User
	conn := r.db.GetDB(ctx)

	if err := conn.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, xerrors.Errorf("failed to SQL execution: %w", err)
	}

	return user.ToModel(), nil
}
