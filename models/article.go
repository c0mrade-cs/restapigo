package models

import "time"

type Content struct {
	Title string
	Body  string
}

type Article struct {
	ID string
	Content
	AuthorID  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
type Articlecreate struct {
	Content
	AuthorID string
}
type Articleupdate struct {
	ID string
	Content
	AuthorID string
}
