package routes

import (
	"golang-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, articleHandler *handler.ArticleHandler) {
	api := r.Group("/api")
	{
		api.POST("/articles", articleHandler.CreateArticle)
		api.GET("/articles", articleHandler.GetAllArticles)
		api.GET("/articles/:id", articleHandler.GetArticleByID)
		api.PUT("/articles/:id", articleHandler.UpdateArticle)
		api.DELETE("/articles/:id", articleHandler.DeleteArticle)
	}
}
