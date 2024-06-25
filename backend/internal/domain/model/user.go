package model

import "time"

type User struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Password  string     `json:"-"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

func NewUser(name, password string) *User {
	t := time.Now()
	return &User{
		Name:      name,
		Password:  password,
		CreatedAt: t,
		UpdatedAt: t,
	}
}

type UserSigninParam struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
