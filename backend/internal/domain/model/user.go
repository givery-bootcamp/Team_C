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
