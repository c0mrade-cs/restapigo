package models

import "time"

// Author ...
type Author struct {
	ID        string
	Firstname string
	Lastname  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// AuthorCreate ...
type AuthorCreate struct {
	Firstname string
	Lastname  string
}

// AuthorUpdate ...
type AuthorUpdate struct {
	ID        string
	Firstname string
	Lastname  string
}

// PackedArticleModel ...
type PackedArticleModel struct {
	ID string
	Content
	Author    Author
	CreatedAt time.Time
	UpdatedAt *time.Time
}
