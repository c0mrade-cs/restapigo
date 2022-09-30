package main

import (
	"net/http"
	"time"

	_ "github.com/c0mrade-cs/article/docs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

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

var InMemoryArticleData []Article

func main() {
	InMemoryArticleData = make([]Article, 0)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/article", CreateArticle)
	r.GET("/article", ReadArticle)
	r.GET("/article/:id", ReadbyIdArticle)
	r.PUT("/article", UpdateArticle)
	r.DELETE("/article/:id", DeleteArticle)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// CreateArticle godoc
// @Summary      Create article
// @Description  create article
// @Tags         article
// @Accept       json
// @Produce      json
// @Param        article    body     Articlecreate  true  "article body"
// @Success      200  {object}   JSONResponse{data=[]Article}
// @Failure      400  {object}  JSONErrorResponse
// @Router       /article [post]
func CreateArticle(c *gin.Context) {
	var data Article
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, JSONErrorResponse{Error: err.Error()})
		return
	}
	data.ID = uuid.New().String()
	t := time.Now()
	data.CreatedAt = &t
	InMemoryArticleData = append(InMemoryArticleData, data)
	c.JSON(http.StatusOK, JSONResponse{
		Data: InMemoryArticleData,
	})
}

// ReadArticle godoc
// @Summary      List article
// @Description  get article
// @Tags         article
// @Accept       json
// @Produce      json
// @Success      200  {object}   JSONResponse{data=[]Article}
// @Router       /article [get]
func ReadArticle(c *gin.Context) {
	c.JSON(http.StatusOK, JSONResponse{
		Data: InMemoryArticleData,
	})
}

// ReadbyIdArticle godoc
// @Summary      get article by id
// @Description  get article by id
// @Tags         article
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "article id"
// @Success      200  {object}   JSONResponse{data=Article}
// @Failure      400  {object}  JSONErrorResponse
// @Router       /article/{id} [get]
func ReadbyIdArticle(c *gin.Context) {
	id := c.Param("id")
	for _, v := range InMemoryArticleData {
		if v.ID == id {
			c.JSON(http.StatusOK, JSONResponse{
				Data: v,
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, JSONErrorResponse{
		Error: "Not Found",
	})

}

// UpdateArticle godoc
// @Summary      Update article
// @Description  update article
// @Tags         article
// @Accept       json
// @Produce      json
// @Param        article    body     Articleupdate  true  "article body"
// @Success      200  {object}   JSONResponse{data=[]Article}
// @Failure      400  {object}  JSONErrorResponse
// @Failure      404  {object}  JSONErrorResponse
// @Router       /article [put]
func UpdateArticle(c *gin.Context) {
	var data Article
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, JSONErrorResponse{Error: err.Error()})
		return
	}
	for i, v := range InMemoryArticleData {
		if v.ID == data.ID {
			t := time.Now()
			data.UpdatedAt = &t
			data.CreatedAt = v.CreatedAt
			InMemoryArticleData[i] = data
			c.JSON(http.StatusOK, JSONResponse{
				Data: InMemoryArticleData,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, JSONErrorResponse{
		Error: "Not Found",
	})
}

// DeleteArticle godoc
// @Summary      Delete article by id
// @Description  delete article by id
// @Tags         article
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "article id"
// @Success      200  {object}   JSONResponse{data=Article}
// @Failure      404  {object}  JSONErrorResponse
// @Router       /article/{id} [delete]
func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	for i, v := range InMemoryArticleData {
		if v.ID == id {
			InMemoryArticleData = remove(InMemoryArticleData, i)
			c.JSON(http.StatusOK, JSONResponse{
				Data: v,
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, JSONErrorResponse{
		Error: "Not Found",
	})
}
func remove(slice []Article, s int) []Article {
	return append(slice[:s], slice[s+1:]...)
}
