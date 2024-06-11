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

func NewPost(title string, body string, user User) *Post {
	return &Post{
		Title:     title,
		Body:      body,
		User:      user,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type CreatePostParam struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func UpdatePost(post *Post,title string, body string) *Post {
  post.Title = title
  post.Body = body
  post.UpdatedAt = time.Now()
  return post
}

type UpdatePostParam struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
