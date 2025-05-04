package routes

import (
	"golang-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupArticleRoutes(api *gin.RouterGroup, articleHandler *handler.ArticleHandler) {
	articles := api.Group("/articles")
	{
		articles.POST("", articleHandler.CreateArticle)
		articles.GET("", articleHandler.GetAllArticles)
		articles.GET("/:id", articleHandler.GetArticleByID)
		articles.PUT("/:id", articleHandler.UpdateArticle)
		articles.DELETE("/:id", articleHandler.DeleteArticle)
	}
}
