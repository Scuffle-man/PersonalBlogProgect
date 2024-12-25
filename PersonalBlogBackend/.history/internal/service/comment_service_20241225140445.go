package service

import (
	"time"

	"gorm.io/gorm"
	"PersonalBlogBackend/internal/models"
)

type CommentService struct {
	db *gorm.DB
}

func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{db: db}
}

// CreateComment 创建评论
func (s *CommentService) CreateComment(content string, articleID, userID uint) (*models.Comment, error) {
	comment := &models.Comment{
		Content:   content,
		ArticleID: articleID,
		UserID:    userID,
		CreatedAt: time.Now(),
	}

	if err := s.db.Create(comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}

// GetCommentsByArticleID 获取文章的所有评论
func (s *CommentService) GetCommentsByArticleID(articleID uint) ([]models.Comment, error) {
	var comments []models.Comment
	if err := s.db.Where("article_id = ?", articleID).
		Preload("User").
		Order("created_at desc").
		Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
} 