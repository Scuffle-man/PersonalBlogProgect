package service

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"PersonalBlogBackend/internal/models"
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

// UpdateArticle 更新文章
func (s *ArticleService) UpdateArticle(id uint, title, content string, userID uint) (*models.Article, error) {
	var article models.Article
	if err := s.db.First(&article, id).Error; err != nil {
		return nil, err
	}

	// 检查是否是文章作者
	if article.AuthorID != userID {
		return nil, errors.New("没有权限修改此文章")
	}

	article.Title = title
	article.Content = content
	article.UpdatedAt = time.Now()

	if err := s.db.Save(&article).Error; err != nil {
		return nil, err
	}

	return &article, nil
}

// DeleteArticle 删除文章
func (s *ArticleService) DeleteArticle(id, userID uint) error {
	var article models.Article
	if err := s.db.First(&article, id).Error; err != nil {
		return err
	}

	// 检查是否是文章作者
	if article.AuthorID != userID {
		return errors.New("没有权限删除此文章")
	}

	// 删除文章及其关联的评论
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 删除文章的所有评论
		if err := tx.Where("article_id = ?", id).Delete(&models.Comment{}).Error; err != nil {
			return err
		}
		// 删除文章
		if err := tx.Delete(&article).Error; err != nil {
			return err
		}
		return nil
	})
} 