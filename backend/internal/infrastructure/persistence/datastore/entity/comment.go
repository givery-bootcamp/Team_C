package entity

import "time"

type Comment struct {
	ID        int
	Body      string
	UserID    int
	User      User `gorm:"foreignKey:UserID"`
	PostID    int
	Post      Post `gorm:"foreignKey:PostID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
