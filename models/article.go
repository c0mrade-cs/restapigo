package models

import "time"

type Person struct {
	Firstname string
	Lastname  string
}

type Content struct {
	Title string
	Body  string
}

type Article struct {
	ID string
	Content
	Author    Person
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
type Articlecreate struct {
	Content
	Author Person
}
type Articleupdate struct {
	ID string
	Content
	Author Person
}
type JSONResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type JSONErrorResponse struct {
	Error string `json:"error"`
}
