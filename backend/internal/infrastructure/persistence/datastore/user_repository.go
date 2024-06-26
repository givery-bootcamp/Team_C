package datastore

import (
	"context"
	"errors"

	"myapp/internal/domain/model"
	"myapp/internal/domain/repository"
	"myapp/internal/exception"
	"myapp/internal/infrastructure/persistence/datastore/driver"
	"myapp/internal/infrastructure/persistence/datastore/entity"

	"github.com/go-sql-driver/mysql"
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

func (r *UserRepository) GetByName(ctx context.Context, name string) (*model.User, error) {
	var user entity.User
	conn := r.db.GetDB(ctx)
	if err := conn.
		Where("name = ?", name).
		First(&user).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, xerrors.Errorf("failed to get sign in params: %w", exception.RecordNotFoundError)
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

func (r *UserRepository) Create(ctx context.Context, user model.User) (*model.User, error) {
	conn := r.db.GetDB(ctx)
	u := entity.NewUserFromModel(&user)

	if err := conn.Create(&u).Error; err != nil {
		mysqlErr := err.(*mysql.MySQLError)
		switch mysqlErr.Number {
		case 1062:
			return nil, xerrors.Errorf("failed to create user: %w", exception.UserAlreadyExistsError)
		}
		return nil, xerrors.Errorf("failed to SQL execution: %w", err)
	}

	return u.ToModel(), nil
}
