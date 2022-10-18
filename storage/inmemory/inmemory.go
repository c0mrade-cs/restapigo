package inmemory

import "github.com/c0mrade-cs/article/models"

// InMemory ...
type InMemory struct {
	Db *DB
}

// DB Mock
type DB struct {
	// InMemoryArticleData ...
	InMemoryArticleData []models.Article
	// InMemoryAuthorData ...
	InMemoryAuthorData []models.Author
}
