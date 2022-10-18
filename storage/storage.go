package storage

import "github.com/c0mrade-cs/article/models"

// StorageI ...
type StorageI interface {
	CreateArticle(id string, entity models.ArticleCreate) error
	ReadbyIDArticle(id string) (models.PackedArticleModel, error)
	ReadArticle(offset, limit int, search string) (resp []models.Article, err error)
	UpdateArticle(entity models.ArticleUpdate) error
	DeleteArticle(id string) error

	CreateAuthor(id string, entity models.AuthorCreate) error
	ReadbyIDAuthor(id string) (models.Author, error)
	ReadAuthor(offset, limit int, search string) (resp []models.Author, err error)
	UpdateAuthor(entity models.AuthorUpdate) error
	DeleteAuthor(id string) error
}
