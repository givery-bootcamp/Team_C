package entity

import (
	"myapp/internal/domain/model"
	"time"
)

type Post struct {
	ID        int
	Title     string
	Body      string
	UserID    int
	User      User `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (p Post) ToModel() *model.Post {
	return &model.Post{
		ID:    p.ID,
		Title: p.Title,
		Body:  p.Body,
		User:  *p.User.ToModel(),
	}
}

func ToPostModelListFromEntity(l []*Post) []*model.Post {
	res := []*model.Post{}
	for _, p := range l {
		res = append(res, p.ToModel())
	}

	return res
}
