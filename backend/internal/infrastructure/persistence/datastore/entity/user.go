package entity

import (
	"myapp/internal/domain/model"
	"time"
)

type User struct {
	ID        int
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (u User) ToModel() *model.User {
	return &model.User{
		ID:        u.ID,
		Name:      u.Name,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}
}
