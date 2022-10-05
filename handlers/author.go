package handlers

import (
	"net/http"

	"github.com/c0mrade-cs/article/models"
	"github.com/c0mrade-cs/article/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateAuthor godoc
// @Summary      Create author
// @Description  create author
// @Tags         author
// @Accept       json
// @Produce      json
// @Param        author    body     models.Authorcreate  true  "author body"
// @Success      201  {object}   models.JSONResponse{data=models.Author}
// @Failure      400  {object}  models.JSONErrorResponse
// @Router       /author [post]
func CreateAuthor(c *gin.Context) {
	var data models.Authorcreate
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	id := uuid.New()
	err := storage.CreateAuthor(id.String(), data)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	author, err := storage.ReadbyIdAuthor(id.String())
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
// @Summary      List author
// @Description  get author
// @Tags         author
// @Accept       json
// @Produce      json
// @Success      200  {object}   models.JSONResponse{data=[]models.Author}
// @Router       /author [get]
func ReadAuthor(c *gin.Context) {
	authorList, err := storage.ReadAuthor()
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

// ReadbyIdAuthor godoc
// @Summary      get author by id
// @Description  get author by id
// @Tags         author
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "author id"
// @Success      200  {object}   models.JSONResponse{data=models.Author}
// @Failure      400  {object}  models.JSONErrorResponse
// @Router       /author/{id} [get]
func ReadbyIdAuthor(c *gin.Context) {
	idStr := c.Param("id")
	author, err := storage.ReadbyIdAuthor(idStr)
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
// @Summary      Update author
// @Description  update author
// @Tags         author
// @Accept       json
// @Produce      json
// @Param        author    body     models.Authorupdate  true  "author body"
// @Success      200  {object}   models.JSONResponse{data=models.Author}
// @Failure      400  {object}  models.JSONErrorResponse
// @Failure      404  {object}  models.JSONErrorResponse
// @Router       /author [put]
func UpdateAuthor(c *gin.Context) {
	var data models.Authorupdate
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	err := storage.UpdateAuthor(data)
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.JSONResponse{
		Data: storage.InMemoryAuthorData,
	})

}

// DeleteAuthor godoc
// @Summary      Delete author by id
// @Description  delete author by id
// @Tags         author
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "author id"
// @Success      200  {object}   models.JSONResponse{data=models.Author}
// @Failure      404  {object}  models.JSONErrorResponse
// @Router       /author/{id} [delete]
func DeleteAuthor(c *gin.Context) {
	id := c.Param("id")
	author, err := storage.DeleteAuthori(id)
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
