package main

import (
	"net/http"

	_ "github.com/c0mrade-cs/article/docs"
	"github.com/c0mrade-cs/article/handlers"
	"github.com/c0mrade-cs/article/models"
	"github.com/c0mrade-cs/article/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func main() {
	storage.InMemoryArticleData = make([]models.Article, 0)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/article", handlers.CreateArticle)
	r.GET("/article", handlers.ReadArticle)
	r.GET("/article/:id", handlers.ReadbyIdArticle)
	r.PUT("/article", handlers.UpdateArticle)
	r.DELETE("/article/:id", handlers.DeleteArticle)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
