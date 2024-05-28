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

<<<<<<< HEAD
func NewPostFromModel(p *model.Post) *Post {
=======
func NewFromModel(p *model.Post) *Post {
>>>>>>> f71fd74 (add: 記事投稿API)
	return &Post{
		ID:        p.ID,
		Title:     p.Title,
		Body:      p.Body,
		UserID:    p.User.ID,
<<<<<<< HEAD
		User:      *NewUserFromModel(&p.User),
=======
>>>>>>> f71fd74 (add: 記事投稿API)
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		DeletedAt: p.DeletedAt,
	}
}

func (p Post) ToModel() *model.Post {
	return &model.Post{
		ID:        p.ID,
		Title:     p.Title,
		Body:      p.Body,
		User:      *p.User.ToModel(),
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		DeletedAt: p.DeletedAt,
	}
}

func ToPostModelListFromEntity(l []*Post) []*model.Post {
	res := []*model.Post{}
	for _, p := range l {
		res = append(res, p.ToModel())
	}

	return res
}
