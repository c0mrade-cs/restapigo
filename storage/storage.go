package storage

import (
	"errors"
	"time"

	"github.com/c0mrade-cs/article/models"
)

var InMemoryArticleData []models.Article

func CreateArticle(id string, entity models.Articlecreate) error {
	var article models.Article
	article.ID = id
	article.Content = entity.Content
	article.AuthorID = entity.AuthorID

	t := time.Now()
	article.CreatedAt = &t
	InMemoryArticleData = append(InMemoryArticleData, article)
	return nil
}

func ReadbyIdArticle(id string) (models.PackedArticleModel, error) {
	var result models.PackedArticleModel
	for _, v := range InMemoryArticleData {
		if v.ID == id {
			author, err := ReadbyIdAuthor(v.AuthorID)
			if err != nil {
				return result, err
			}

			result.ID = v.ID
			result.Content = v.Content
			result.Author = author
			result.CreatedAt = *v.CreatedAt
			result.UpdatedAt = v.UpdatedAt
			return result, nil
		}
	}
	return models.PackedArticleModel{}, errors.New("article not found")
}

func ReadArticle() (resp []models.Article, err error) {
	resp = InMemoryArticleData
	return resp, err
}

func UpdateArticle(entity models.Articleupdate) error {
	var article models.Article
	for i, v := range InMemoryArticleData {
		if v.ID == entity.ID {
			t := time.Now()
			article.UpdatedAt = &t
			article.CreatedAt = v.CreatedAt

			article.Content = entity.Content
			article.AuthorID = entity.AuthorID
			InMemoryArticleData[i] = article

		}
	}

	return nil
}

func DeleteArticlei(id string) (models.Article, error) {
	for i, v := range InMemoryArticleData {
		if v.ID == id {
			InMemoryArticleData = remove(InMemoryArticleData, i)

			return v, nil
		}
	}
	return models.Article{}, errors.New("article not found")
}

func remove(slice []models.Article, s int) []models.Article {
	return append(slice[:s], slice[s+1:]...)
}
