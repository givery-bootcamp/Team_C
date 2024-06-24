package datastore

import (
	"context"

	"myapp/internal/domain/model"
	"myapp/internal/domain/repository"
	"myapp/internal/exception"

	"myapp/internal/infrastructure/persistence/datastore/driver"
	"myapp/internal/infrastructure/persistence/datastore/entity"

	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

type PostRepository struct {
	db driver.DB
}

func NewPostRepository(db driver.DB) repository.PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (r *PostRepository) GetAll(ctx context.Context, limit, offset int) ([]*model.Post, error) {
	posts := []*entity.Post{}

	conn := r.db.GetDB(ctx)
	if err := conn.Preload("User").Limit(limit).Offset(offset).Find(&posts).Error; err != nil {
		return nil, xerrors.Errorf("failed to SQL execution: %w", err)
	}
	return entity.ToPostModelListFromEntity(posts), nil
}

func (r *PostRepository) GetByID(ctx context.Context, id int) (*model.Post, error) {
	var p entity.Post

	conn := r.db.GetDB(ctx)
	if err := conn.Preload("User").Where("id = ?", id).First(&p).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.RecordNotFoundError
		}
		return nil, xerrors.Errorf("failed to SQL execution: %w", err)
	}

	return p.ToModel(), nil
}

func (r *PostRepository) Create(ctx context.Context, post *model.Post) (*model.Post, error) {
	p := entity.NewPostFromModel(post)

	conn := r.db.GetDB(ctx)

	res := conn.Create(&p)
	if res.Error != nil {
		return nil, xerrors.Errorf("failed to SQL execution: %w", res.Error)
	}

	user := entity.User{}
	if err := conn.Where("id = ?", p.UserID).First(&user).Error; err != nil {
		return nil, xerrors.Errorf("failed to SQL execution: %w", err)
	}
	p.User = user

	return p.ToModel(), nil
}

func (r *PostRepository) Update(ctx context.Context, post *model.Post) (*model.Post, error) {
	p := entity.NewPostFromModel(post)

	conn := r.db.GetDB(ctx)

	res := conn.Model(&entity.Post{}).
		Where("id = ?", p.ID).
		Updates(p)

	if res.Error != nil {
		return nil, res.Error
	}

	return p.ToModel(), nil
}

func (r *PostRepository) Delete(ctx context.Context, postId int) error {
	conn := r.db.GetDB(ctx)

	if err := conn.Where("post_id = ?", postId).Delete(&entity.Comment{}).Error; err != nil {
		return xerrors.Errorf("failed to delete comments: %w", err)
	}

	res := conn.Where("id = ?", postId).Delete(&entity.Post{})
	if res.Error != nil {
		return xerrors.Errorf("failed to execute SQL: %w", res.Error)
	}

	return nil
}
