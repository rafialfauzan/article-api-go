// internal/handler/article_handler.go
package handler

import (
	"golang-api/internal/model"
	"golang-api/internal/service"
	"golang-api/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	service service.ArticleService
}

func NewArticleHandler(service service.ArticleService) *ArticleHandler {
	return &ArticleHandler{service}
}

func (h *ArticleHandler) CreateArticle(c *gin.Context) {
	var article model.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		utils.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.CreateArticle(&article); err != nil {
		utils.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONSuccess(c, http.StatusCreated, "Article created successfully", article)
}

func (h *ArticleHandler) GetAllArticles(c *gin.Context) {
	articles, err := h.service.GetAllArticles()
	if err != nil {
		utils.JSONError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONSuccess(c, http.StatusOK, "Articles retrieved successfully", articles)
}

func (h *ArticleHandler) GetArticleByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.JSONError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	article, err := h.service.GetArticleByID(uint(id))
	if err != nil {
		utils.JSONError(c, http.StatusNotFound, err.Error())
		return
	}

	utils.JSONSuccess(c, http.StatusOK, "Article retrieved successfully", article)
}

func (h *ArticleHandler) UpdateArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.JSONError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var article model.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		utils.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.UpdateArticle(uint(id), &article); err != nil {
		utils.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONSuccess(c, http.StatusOK, "Article updated successfully", article)
}

func (h *ArticleHandler) DeleteArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.JSONError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := h.service.DeleteArticle(uint(id)); err != nil {
		utils.JSONError(c, http.StatusNotFound, err.Error())
		return
	}

	utils.JSONSuccess(c, http.StatusOK, "Article deleted successfully", nil)
}
