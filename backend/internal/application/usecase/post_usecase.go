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

func (u *PostUsecase) Create(ctx context.Context, title string, body string) (*model.Post, error) {
	// TODO: cookieからユーザ情報を取得する

	post := model.NewPost(title, body, model.User{
		// TODO: データベースから取得する
		ID:   1,
		Name: "taro",
	})

	return u.r.Create(ctx, post)
}
