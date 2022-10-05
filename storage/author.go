package storage

import (
	"errors"
	"time"

	"github.com/c0mrade-cs/article/models"
)

var InMemoryAuthorData []models.Author

func CreateAuthor(id string, entity models.Authorcreate) error {
	var author models.Author
	author.ID = id
	author.Firstname = entity.Firstname
	author.Lastname = entity.Lastname

	t := time.Now()
	author.CreatedAt = &t
	InMemoryAuthorData = append(InMemoryAuthorData, author)
	return nil
}

func ReadbyIdAuthor(id string) (models.Author, error) {
	for _, v := range InMemoryAuthorData {
		if v.ID == id {
			return v, nil
		}
	}
	return models.Author{}, errors.New("author not found")
}

func ReadAuthor() (resp []models.Author, err error) {
	resp = InMemoryAuthorData
	return resp, err
}

func UpdateAuthor(entity models.Authorupdate) error {
	var author models.Author
	for i, v := range InMemoryAuthorData {
		if v.ID == entity.ID {
			t := time.Now()
			author.UpdatedAt = &t
			author.CreatedAt = v.CreatedAt

			author.Firstname = entity.Firstname
			author.Lastname = entity.Lastname
			InMemoryAuthorData[i] = author

		}
	}

	return nil
}

func DeleteAuthori(id string) (models.Author, error) {
	for i, v := range InMemoryAuthorData {
		if v.ID == id {
			InMemoryAuthorData = removeAuthor(InMemoryAuthorData, i)

			return v, nil
		}
	}
	return models.Author{}, errors.New("author not found")
}

func removeAuthor(slice []models.Author, s int) []models.Author {
	return append(slice[:s], slice[s+1:]...)
}
