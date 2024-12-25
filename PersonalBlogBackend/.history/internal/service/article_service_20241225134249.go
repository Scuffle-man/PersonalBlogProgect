package service

import (
	"time"

	"gorm.io/gorm"
	"your-project/internal/models"
)

type ArticleService struct {
	db *gorm.DB
}

func NewArticleService(db *gorm.DB) *ArticleService {
	return &ArticleService{db: db}
}

// CreateArticle 创建新文章
func (s *ArticleService) CreateArticle(title, content string, authorID uint) (*models.Article, error) {
	article := &models.Article{
		Title:       title,
		Content:     content,
		AuthorID:    authorID,
		PublishedAt: time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.db.Create(article).Error; err != nil {
		return nil, err
	}

	return article, nil
}

// GetArticles 获取文章列表
func (s *ArticleService) GetArticles(page, pageSize int) ([]models.Article, error) {
	var articles []models.Article
	offset := (page - 1) * pageSize

	if err := s.db.Offset(offset).Limit(pageSize).
		Preload("Author").
		Order("published_at desc").
		Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}

// GetArticleByID 获取指定文章
func (s *ArticleService) GetArticleByID(id uint) (*models.Article, error) {
	var article models.Article
	if err := s.db.Preload("Author").Preload("Comments.User").First(&article, id).Error; err != nil {
		return nil, err
	}
	return &article, nil
} 