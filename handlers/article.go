package handlers

import (
	"net/http"
	"strconv"

	"uacademy/blogpost/article_service/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateArticle godoc
// @Summary     Create article
// @Description create a new article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       article body     models.CreateArticleModel true "article body"
// @Success     201     {object} models.JSONResponse{data=models.Article}
// @Failure     400     {object} models.JSONErrorResponse
// @Router      /v1/article [post]
func (h handler) CreateArticle(c *gin.Context) {
	var body models.CreateArticleModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	// TODO - validation should be here

	id := uuid.New()

	err := h.Stg.AddArticle(id.String(), body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	article, err := h.Stg.GetArticleByID(id.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Article | GetList",
		Data:    article,
	})
}

// GetArticleByID godoc
// @Summary     get article by id
// @Description get an article by id
// @Tags        articles
// @Accept      json
// @Param       id path string true "Article ID"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.PackedArticleModel}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v1/article/{id} [get]
func (h handler) GetArticleByID(c *gin.Context) {
	idStr := c.Param("id")

	// TODO - validation
	article, err := h.Stg.GetArticleByID(idStr)
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    article,
	})
}

// GetArticleList godoc
// @Summary     List articles
// @Description get articles
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       offset query    int false "0"
// @Param       limit  query    int false "10"
// @Param       search query    string false "smth"
// @Success     200    {object} models.JSONResponse{data=[]models.Article}
// @Router      /v1/article [get]
func (h handler) GetArticleList(c *gin.Context) {
	offsetStr := c.DefaultQuery("offset", h.Cfg.DefaultOffset)
	limitStr := c.DefaultQuery("limit", h.Cfg.DefaultLimit)
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
	articleList, err := h.Stg.GetArticleList(offset, limit, searchStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    articleList,
	})
}

// UpdateArticle godoc
// @Summary     Update article
// @Description update a new article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       article body     models.UpdateArticleModel true "article body"
// @Success     200     {object} models.JSONResponse{data=models.Article}
// @Failure     400     {object} models.JSONErrorResponse
// @Router      /v1/article [put]
func (h handler) UpdateArticle(c *gin.Context) {
	var body models.UpdateArticleModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.Stg.UpdateArticle(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	article, err := h.Stg.GetArticleByID(body.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    article,
	})
}

// DeleteArticle godoc
// @Summary     delete article by id
// @Description delete an article by id
// @Tags        articles
// @Accept      json
// @Param       id path string true "Article ID"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.PackedArticleModel}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v1/article/{id} [delete]
func (h handler) DeleteArticle(c *gin.Context) {
	idStr := c.Param("id")

	article, err := h.Stg.GetArticleByID(idStr)
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
		Message: "OK",
		Data:    article,
	})
}
