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

func (u *PostUsecase) GetAll(ctx context.Context, limit, offset int) ([]*model.Post, error) {
	return u.r.GetAll(ctx, limit, offset)
}

func (u *PostUsecase) GetByID(ctx context.Context, id int) (*model.Post, error) {
	return u.r.GetByID(ctx, id)
}

<<<<<<< HEAD
<<<<<<< HEAD
func (u *PostUsecase) Create(ctx context.Context, title, body string, userId int) (*model.Post, error) {
	post := model.NewPost(title, body, model.User{
		ID: userId,
=======
func (u *PostUsecase) Create(ctx context.Context, title string, body string) (*model.Post, error) {
	// TODO: cookieからユーザ情報を取得する

	post := model.NewPost(title, body, model.User{
		// TODO: データベースから取得する
		ID:   1,
		Name: "taro",
>>>>>>> f71fd74 (add: 記事投稿API)
=======
func (u *PostUsecase) Create(ctx context.Context, title string, body string, userId int) (*model.Post, error) {
	post := model.NewPost(title, body, model.User{
		ID: userId,
>>>>>>> aadcb65 (update: cookieからUserId取得して仮実装削除)
	})

	return u.r.Create(ctx, post)
}
