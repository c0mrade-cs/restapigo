package main

import (
	"net/http"

	_ "github.com/c0mrade-cs/article/docs"
	"github.com/c0mrade-cs/article/handlers"
	"github.com/c0mrade-cs/article/storage"
	"github.com/c0mrade-cs/article/storage/inmemory"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	var stg storage.StorageI
	stg = inmemory.InMemory{
		Db: &inmemory.DB{},
	}
	h := handlers.Handler{
		Stg: stg,
	}

	r.POST("/article", h.CreateArticle)
	r.GET("/article", h.ReadArticle)
	r.GET("/article/:id", h.ReadbyIDArticle)
	r.PUT("/article", h.UpdateArticle)
	r.DELETE("/article/:id", h.DeleteArticle)

	r.POST("/author", h.CreateAuthor)
	r.GET("/author", h.ReadAuthor)
	r.GET("/author/:id", h.ReadbyIDAuthor)
	r.PUT("/author", h.UpdateAuthor)
	r.DELETE("/author/:id", h.DeleteAuthor)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
