package handlers

import (
	"net/http"
	"strconv"

	"github.com/c0mrade-cs/article/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateAuthor godoc
// @Summary     Create author
// @Description create author
// @Tags        author
// @Accept      json
// @Produce     json
// @Param       author body     models.AuthorCreate true "author body"
// @Success     201    {object} models.JSONResponse{data=models.Author}
// @Failure     400    {object} models.JSONErrorResponse
// @Router      /author [post]
func (h Handler) CreateAuthor(c *gin.Context) {
	var data models.AuthorCreate
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	id := uuid.New()
	err := h.Stg.CreateAuthor(id.String(), data)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	author, err := h.Stg.ReadbyIDAuthor(id.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResponse{
		Data: author,
	})
}

// ReadAuthor godoc
// @Summary     List author
// @Description get author
// @Tags        author
// @Accept      json
// @Produce     json
// @Param       offset query    int    false "0"
// @Param       limit  query    int    false "10"
// @Param       search query    string false "smth"
// @Success     200 {object} models.JSONResponse{data=[]models.Author}
// @Router      /author [get]
func (h Handler) ReadAuthor(c *gin.Context) {
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
	authorList, err := h.Stg.ReadAuthor(offset, limit, searchStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Data: authorList,
	})
}

// ReadbyIDAuthor godoc
// @Summary     get author by id
// @Description get author by id
// @Tags        author
// @Accept      json
// @Produce     json
// @Param       id  path     string true "author id"
// @Success     200 {object} models.JSONResponse{data=models.Author}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /author/{id} [get]
func (h Handler) ReadbyIDAuthor(c *gin.Context) {
	idStr := c.Param("id")

	author, err := h.Stg.ReadbyIDAuthor(idStr)
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Data: author,
	})
}

// UpdateAuthor godoc
// @Summary     Update author
// @Description update author
// @Tags        author
// @Accept      json
// @Produce     json
// @Param       author body     models.AuthorUpdate true "author body"
// @Success     200    {object} models.JSONResponse{data=models.Author}
// @Failure     400    {object} models.JSONErrorResponse
// @Failure     404    {object} models.JSONErrorResponse
// @Router      /author [put]
func (h Handler) UpdateAuthor(c *gin.Context) {
	var data models.AuthorUpdate
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	err := h.Stg.UpdateAuthor(data)
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	author, err := h.Stg.ReadbyIDAuthor(data.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Data: author,
	})
}

// DeleteAuthor godoc
// @Summary     Delete author by id
// @Description delete author by id
// @Tags        author
// @Accept      json
// @Produce     json
// @Param       id  path     string true "author id"
// @Success     200 {object} models.JSONResponse{data=models.Author}
// @Failure     404 {object} models.JSONErrorResponse
// @Router      /author/{id} [delete]
func (h Handler) DeleteAuthor(c *gin.Context) {
	id := c.Param("id")

	author, err := h.Stg.ReadbyIDAuthor(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	err = h.Stg.DeleteAuthor(author.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Data: author,
	})
}
