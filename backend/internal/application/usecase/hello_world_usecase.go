package usecase

import (
	"context"
	"myapp/internal/domain/model"
	"myapp/internal/domain/repository"
)

type HelloWorldUsecase struct {
	r repository.HelloWorldRepository
}

func NewHelloWorldUsecase(r repository.HelloWorldRepository) HelloWorldUsecase {
	return HelloWorldUsecase{
		r: r,
	}
}
func (u *HelloWorldUsecase) Execute(ctx context.Context, lang string) (*model.HelloWorld, error) {
	return u.r.Get(ctx, lang)
}
