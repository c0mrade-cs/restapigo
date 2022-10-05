package handlers

import (
	"net/http"

	"github.com/c0mrade-cs/article/models"
	"github.com/c0mrade-cs/article/storage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateArticle godoc
// @Summary      Create article
// @Description  create article
// @Tags         article
// @Accept       json
// @Produce      json
// @Param        article    body     models.ArticleCreate  true  "article body"
// @Success      201  {object}   models.JSONResponse{data=models.Article}
// @Failure      400  {object}  models.JSONErrorResponse
// @Router       /article [post]
func CreateArticle(c *gin.Context) {
	var data models.ArticleCreate
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	id := uuid.New()
	err := storage.CreateArticle(id.String(), data)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	article, err := storage.ReadbyIDArticle(id.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResponse{
		Data: article,
	})
}

// ReadArticle godoc
// @Summary      List article
// @Description  get article
// @Tags         article
// @Accept       json
// @Produce      json
// @Success      200  {object}   models.JSONResponse{data=[]models.Article}
// @Router       /article [get]
func ReadArticle(c *gin.Context) {
	articleList, err := storage.ReadArticle()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Data: articleList,
	})
}

// ReadbyIDArticle godoc
// @Summary      get article by id
// @Description  get article by id
// @Tags         article
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "article id"
// @Success      200  {object}   models.JSONResponse{data=models.PackedArticleModel}
// @Failure      400  {object}  models.JSONErrorResponse
// @Router       /article/{id} [get]
func ReadbyIDArticle(c *gin.Context) {
	idStr := c.Param("id")
	article, err := storage.ReadbyIDArticle(idStr)
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Data: article,
	})

}

// UpdateArticle godoc
// @Summary      Update article
// @Description  update article
// @Tags         article
// @Accept       json
// @Produce      json
// @Param        article    body     models.ArticleUpdate  true  "article body"
// @Success      200  {object}   models.JSONResponse{data=models.Article}
// @Failure      400  {object}  models.JSONErrorResponse
// @Failure      404  {object}  models.JSONErrorResponse
// @Router       /article [put]
func UpdateArticle(c *gin.Context) {
	var data models.ArticleUpdate
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	err := storage.UpdateArticle(data)
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Data: storage.InMemoryArticleData,
	})

}

// DeleteArticle godoc
// @Summary      Delete article by id
// @Description  delete article by id
// @Tags         article
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "article id"
// @Success      200  {object}   models.JSONResponse{data=models.Article}
// @Failure      404  {object}  models.JSONErrorResponse
// @Router       /article/{id} [delete]
func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	article, err := storage.DeleteArticlei(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Data: article,
	})
}
