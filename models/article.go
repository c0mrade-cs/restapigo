package models

import "time"

// Content ...
type Content struct {
	Title string
	Body  string
}

// Article ...
type Article struct {
	ID string
	Content
	AuthorID  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// ArticleCreate ...
type ArticleCreate struct {
	Content
	AuthorID string
}

// ArticleUpdate ...
type ArticleUpdate struct {
	ID string
	Content
	AuthorID string
}
