package service

import (
	"errors"
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
	// 检查文章是否存在
	var article models.Article
	if err := s.db.First(&article, articleID).Error; err != nil {
		return nil, errors.New("文章不存在")
	}

	comment := &models.Comment{
		Content:   content,
		ArticleID: articleID,
		UserID:    userID,
		CreatedAt: time.Now(),
	}

	if err := s.db.Create(comment).Error; err != nil {
		return nil, err
	}

	// 加载评论用户信息
	if err := s.db.Preload("User").First(comment, comment.ID).Error; err != nil {
		return nil, err
	}

	return comment, nil
}

// GetCommentsByArticleID 获取文章的所有评论
func (s *CommentService) GetCommentsByArticleID(articleID uint) ([]models.Comment, error) {
	// 检查文章是否存在
	var article models.Article
	if err := s.db.First(&article, articleID).Error; err != nil {
		return nil, errors.New("文章不存在")
	}

	var comments []models.Comment
	if err := s.db.Where("article_id = ?", articleID).
		Preload("User").
			Order("created_at desc").
			Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

// DeleteComment 删除评论
func (s *CommentService) DeleteComment(commentID, userID uint) error {
	var comment models.Comment
	if err := s.db.First(&comment, commentID).Error; err != nil {
		return errors.New("评论不存在")
	}

	// 检查是否是评论作者
	if comment.UserID != userID {
		return errors.New("没有权限删除此评论")
	}

	return s.db.Delete(&comment).Error
} 