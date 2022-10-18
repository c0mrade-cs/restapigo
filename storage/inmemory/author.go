package inmemory

import (
	"errors"
	"strings"
	"time"

	"github.com/c0mrade-cs/article/models"
)

// CreateAuthor ...
func (im InMemory) CreateAuthor(id string, entity models.AuthorCreate) error {
	var author models.Author
	author.ID = id
	author.Firstname = entity.Firstname
	author.Lastname = entity.Lastname

	t := time.Now()
	author.CreatedAt = &t
	im.Db.InMemoryAuthorData = append(im.Db.InMemoryAuthorData, author)
	return nil
}

// ReadbyIDAuthor ...
func (im InMemory) ReadbyIDAuthor(id string) (models.Author, error) {
	for _, v := range im.Db.InMemoryAuthorData {
		if v.ID == id && v.DeletedAt == nil {
			return v, nil
		}
	}
	return models.Author{}, errors.New("author not found")
}

// ReadAuthor ...
func (im InMemory) ReadAuthor(offset, limit int, search string) (resp []models.Author, err error) {
	off := 0
	c := 0
	for _, v := range im.Db.InMemoryAuthorData {
		if v.DeletedAt == nil && (strings.Contains(v.Firstname, search) || strings.Contains(v.Lastname, search)) {
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

// UpdateAuthor ...
func (im InMemory) UpdateAuthor(entity models.AuthorUpdate) error {
	for i, v := range im.Db.InMemoryAuthorData {
		if v.ID == entity.ID && v.DeletedAt == nil {
			t := time.Now()
			v.UpdatedAt = &t
			v.Firstname = entity.Firstname
			v.Lastname = entity.Lastname
			im.Db.InMemoryAuthorData[i] = v
		}
	}

	return nil
}

// DeleteAuthor ...
func (im InMemory) DeleteAuthor(id string) error {
	for i, v := range im.Db.InMemoryAuthorData {
		if v.ID == id {
			if v.DeletedAt != nil {
				return errors.New("author already deleted")
			}
			t := time.Now()
			v.DeletedAt = &t
			im.Db.InMemoryAuthorData[i] = v
			return nil
		}
	}
	return errors.New("author not found")
}
