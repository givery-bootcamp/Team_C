package usecase

import (
	"context"

	"myapp/internal/domain/model"
	"myapp/internal/domain/repository"
	"myapp/internal/exception"
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

func (u *PostUsecase) Create(ctx context.Context, title, body string, userId int) (*model.Post, error) {
	post := model.NewPost(title, body, model.User{
		ID: userId,
	})

	return u.r.Create(ctx, post)
}

func (u *PostUsecase) Update(ctx context.Context, title, body string, postId int, userId int) (*model.Post, error) {
	post, err := u.GetByID(ctx, postId)
	if err != nil {
		return nil, err
	}

	if post.User.ID != userId {
		return nil, exception.RecordNotFoundError
	}

	updatedPost := post.UpdatePost(title, body)

	return u.r.Update(ctx, updatedPost)
}

func (u *PostUsecase) Delete(ctx context.Context, postId int, userId int) error {
	post, err := u.GetByID(ctx, postId)
	if err != nil {
		return err
	}

	if post.User.ID != userId {
		return exception.RecordNotFoundError
	}

	return u.r.Delete(ctx, postId)
}
