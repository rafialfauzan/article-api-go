package service

import (
	"errors"
	"golang-api/internal/model"
	"golang-api/internal/repository"
)

type ArticleService interface {
	CreateArticle(article *model.Article) error
	GetAllArticles() ([]model.Article, error)
	GetArticleByID(id uint) (model.Article, error)
	UpdateArticle(id uint, article *model.Article) error
	DeleteArticle(id uint) error
}

type articleService struct {
	repo repository.ArticleRepository
}

func NewArticleService(repo repository.ArticleRepository) ArticleService {
	return &articleService{repo}
}

func (s *articleService) CreateArticle(article *model.Article) error {
	if article.Title == "" {
		return errors.New("title is required")
	}
	if article.Content == "" {
		return errors.New("content is required")
	}
	return s.repo.Create(article)
}

func (s *articleService) GetAllArticles() ([]model.Article, error) {
	return s.repo.FindAll()
}

func (s *articleService) GetArticleByID(id uint) (model.Article, error) {
	article, err := s.repo.FindByID(id)
	if err != nil {
		return model.Article{}, errors.New("article not found")
	}
	return article, nil
}

func (s *articleService) UpdateArticle(id uint, article *model.Article) error {
	if article.Title == "" {
		return errors.New("title is required")
	}
	if article.Content == "" {
		return errors.New("content is required")
	}

	existingArticle, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("article not found")
	}

	existingArticle.Title = article.Title
	existingArticle.Content = article.Content
	return s.repo.Update(&existingArticle)
}

func (s *articleService) DeleteArticle(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("article not found")
	}
	return s.repo.Delete(id)
}
