package repository

import (
	"golang-api/internal/model"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	Create(article *model.Article) error
	FindAll() ([]model.Article, error)
	FindByID(id uint) (model.Article, error)
	Update(article *model.Article) error
	Delete(id uint) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db}
}

func (r *articleRepository) Create(article *model.Article) error {
	return r.db.Create(article).Error
}

func (r *articleRepository) FindAll() ([]model.Article, error) {
	var articles []model.Article
	err := r.db.Find(&articles).Error
	return articles, err
}

func (r *articleRepository) FindByID(id uint) (model.Article, error) {
	var article model.Article
	err := r.db.First(&article, id).Error
	return article, err
}

func (r *articleRepository) Update(article *model.Article) error {
	return r.db.Save(article).Error
}

func (r *articleRepository) Delete(id uint) error {
	return r.db.Delete(&model.Article{}, id).Error
}
