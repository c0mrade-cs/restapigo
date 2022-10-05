package models

import "time"

type Author struct {
	ID        string
	Firstname string
	Lastname  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
type Authorcreate struct {
	Firstname string
	Lastname  string
}
type Authorupdate struct {
	ID        string
	Firstname string
	Lastname  string
}

type PackedArticleModel struct {
	ID string
	Content
	Author    Author
	CreatedAt time.Time
	UpdatedAt *time.Time
}
