package usecase

import (
	"context"
	"myapp/internal/domain/model"
	"myapp/internal/domain/repository"
)

type PostUsecase struct {
	r repository.PostRepository
}

func NewPostUsecase(r repository.PostRepository) PostUsecase {
	return PostUsecase{
		r: r,
	}
}

func (u *PostUsecase) GetAll(ctx context.Context) ([]*model.Post, error) {
	return u.r.GetAll(ctx)
}

func (u *PostUsecase) GetByID(ctx context.Context, id int) (*model.Post, error) {
	return u.r.GetByID(ctx, id)
}

func (u *PostUsecase) Create(ctx context.Context, title, body string, userId int) (*model.Post, error) {
	post := model.NewPost(title, body, model.User{
		ID: userId,
	})

	return u.r.Create(ctx, post)
}
