package handlers

import (
	"net/http"
	"strconv"

	"github.com/c0mrade-cs/article/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateArticle godoc
// @Summary     Create article
// @Description create article
// @Tags        article
// @Accept      json
// @Produce     json
// @Param       article body     models.ArticleCreate true "article body"
// @Success     201     {object} models.JSONResponse{data=models.Article}
// @Failure     400     {object} models.JSONErrorResponse
// @Router      /article [post]
func (h Handler) CreateArticle(c *gin.Context) {
	var data models.ArticleCreate
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	id := uuid.New()
	err := h.Stg.CreateArticle(id.String(), data)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	article, err := h.Stg.ReadbyIDArticle(id.String())
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
// @Summary     List article
// @Description get article
// @Tags        article
// @Accept      json
// @Produce     json
// @Param       offset query    int    false "0"
// @Param       limit  query    int    false "10"
// @Param       search query    string false "smth"
// @Success     200    {object} models.JSONResponse{data=[]models.Article}
// @Router      /article [get]
func (h Handler) ReadArticle(c *gin.Context) {
	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "10")
	searchStr := c.DefaultQuery("search", "")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	articleList, err := h.Stg.ReadArticle(offset, limit, searchStr)
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
// @Summary     get article by id
// @Description get article by id
// @Tags        article
// @Accept      json
// @Produce     json
// @Param       id  path     string true "article id"
// @Success     200 {object} models.JSONResponse{data=models.PackedArticleModel}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /article/{id} [get]
func (h Handler) ReadbyIDArticle(c *gin.Context) {
	idStr := c.Param("id")
	article, err := h.Stg.ReadbyIDArticle(idStr)
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
// @Summary     Update article
// @Description update article
// @Tags        article
// @Accept      json
// @Produce     json
// @Param       article body     models.ArticleUpdate true "article body"
// @Success     200     {object} models.JSONResponse{data=models.Article}
// @Failure     400     {object} models.JSONErrorResponse
// @Failure     404     {object} models.JSONErrorResponse
// @Router      /article [put]
func (h Handler) UpdateArticle(c *gin.Context) {
	var data models.ArticleUpdate
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	err := h.Stg.UpdateArticle(data)
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	article, err := h.Stg.ReadbyIDArticle(data.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Data: article,
	})

}

// DeleteArticle godoc
// @Summary     Delete article by id
// @Description delete article by id
// @Tags        article
// @Accept      json
// @Produce     json
// @Param       id  path     string true "article id"
// @Success     200 {object} models.JSONResponse{data=models.PackedArticleModel}
// @Failure     404 {object} models.JSONErrorResponse
// @Router      /article/{id} [delete]
func (h Handler) DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	article, err := h.Stg.ReadbyIDArticle(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	err = h.Stg.DeleteArticle(article.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Data: article,
	})
}
