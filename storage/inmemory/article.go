package inmemory

import (
	"errors"
	"strings"
	"time"

	"github.com/c0mrade-cs/article/models"
)

// CreateArticle ...
func (im InMemory) CreateArticle(id string, entity models.ArticleCreate) error {
	var article models.Article
	article.ID = id
	article.Content = entity.Content
	author, err := im.ReadbyIDAuthor(entity.AuthorID)
	if err != nil {
		return err
	}
	article.AuthorID = author.ID
	article.CreatedAt = time.Now()
	im.Db.InMemoryArticleData = append(im.Db.InMemoryArticleData, article)
	return nil
}

// ReadbyIDArticle ...
func (im InMemory) ReadbyIDArticle(id string) (models.PackedArticleModel, error) {
	var result models.PackedArticleModel
	for _, v := range im.Db.InMemoryArticleData {
		if v.ID == id && v.DeletedAt == nil {
			author, err := im.ReadbyIDAuthor(v.AuthorID)
			if err != nil {
				return result, err
			}

			result.ID = v.ID
			result.Content = v.Content
			result.Author = author
			result.CreatedAt = v.CreatedAt
			result.UpdatedAt = v.UpdatedAt
			result.DeletedAt = v.DeletedAt
			return result, nil
		}
	}
	return models.PackedArticleModel{}, errors.New("article not found")
}

// ReadArticle ...
func (im InMemory) ReadArticle(offset, limit int, search string) (resp []models.Article, err error) {
	off := 0
	c := 0
	for _, v := range im.Db.InMemoryArticleData {
		if v.DeletedAt == nil && (strings.Contains(v.Title, search) || strings.Contains(v.Body, search)) {
			if offset <= off {
				c++
				resp = append(resp, v)
			}

			if c >= limit {
				break
			}

			off++
		}
	}

	return resp, err
}

// UpdateArticle ...
func (im InMemory) UpdateArticle(entity models.ArticleUpdate) error {

	for i, v := range im.Db.InMemoryArticleData {
		if v.ID == entity.ID && v.DeletedAt == nil {
			t := time.Now()
			v.UpdatedAt = &t
			v.Content = entity.Content
			im.Db.InMemoryArticleData[i] = v
			return nil
		}
	}

	return errors.New("article not found")
}

// DeleteArticle ...
func (im InMemory) DeleteArticle(id string) error {
	for i, v := range im.Db.InMemoryArticleData {
		if v.ID == id {
			if v.DeletedAt != nil {
				return errors.New("article already deleted")
			}
			t := time.Now()
			v.DeletedAt = &t
			im.Db.InMemoryArticleData[i] = v
			return nil
		}
	}
	return errors.New("article not found")
}
