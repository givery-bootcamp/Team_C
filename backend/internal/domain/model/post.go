package model

import "time"

type Post struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	User      User       `json:"user"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}
